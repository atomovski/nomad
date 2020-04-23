package volumewatcher

import (
	"context"
	"fmt"
	"sync"

	log "github.com/hashicorp/go-hclog"
	memdb "github.com/hashicorp/go-memdb"
	multierror "github.com/hashicorp/go-multierror"
	cstructs "github.com/hashicorp/nomad/client/structs"
	"github.com/hashicorp/nomad/nomad/state"
	"github.com/hashicorp/nomad/nomad/structs"
)

// volumeWatcher is used to watch a single volume and trigger the
// scheduler when allocation health transitions.
type volumeWatcher struct {
	// v is the volume being watched
	v *structs.CSIVolume

	// state is the state that is watched for state changes.
	state *state.StateStore

	// updateClaims is the function used to apply claims to raft
	updateClaims updateClaimsFn

	// server interface for CSI client RPCs
	rpc ClientRPC

	logger log.Logger
	ctx    context.Context
	exitFn context.CancelFunc

	// updateCh is triggered when there is an updated volume
	updateCh chan *structs.CSIVolume

	wLock sync.RWMutex
}

// newVolumeWatcher returns a volume watcher that is used to watch
// volumes
func newVolumeWatcher(parent *Watcher, vol *structs.CSIVolume) *volumeWatcher {
	ctx, exitFn := context.WithCancel(parent.ctx)
	w := &volumeWatcher{
		updateCh:     make(chan *structs.CSIVolume, 1),
		updateClaims: parent.updateClaims,
		v:            vol,
		state:        parent.state,
		rpc:          parent.rpc,
		logger:       parent.logger.With("volume_id", vol.ID, "namespace", vol.Namespace),
		ctx:          ctx,
		exitFn:       exitFn,
	}

	// Start the long lived watcher that scans for allocation updates
	go w.watch()

	return w
}

// Notify signals an update to the tracked volume.
func (vw *volumeWatcher) Notify(v *structs.CSIVolume) {
	select {
	case vw.updateCh <- v:
	default:
	}
}

// StopWatch stops watching the volume. This should be called whenever a
// volume's claims are fully reaped or the watcher is no longer needed.
func (vw *volumeWatcher) Stop() {
	vw.exitFn()
}

// watch is the long-running function that watches for changes to a volume.
// Each pass steps the volume's claims through the various states of reaping
// until the volume has no more claims eligible to be reaped.
func (vw *volumeWatcher) watch() {
	for {
		select {
		case <-vw.ctx.Done():
			// TODO(tgross): currently server->client RPC have no cancellation
			// context, so we can't stop the long-runner RPCs gracefully
			return
		case vol := <-vw.updateCh:
			// while we won't make raft writes if we get a stale update,
			// we can still fire extra CSI RPC calls if we don't check this
			if vol.ModifyIndex >= vw.v.ModifyIndex {
				vol = vw.getVolume(vol)
				if vol == nil {
					return
				}
				vw.volumeReap(vol)
			}
		}
	}
}

// getVolume returns the tracked volume, fully populated with the current
// state
func (vw *volumeWatcher) getVolume(vol *structs.CSIVolume) *structs.CSIVolume {
	vw.wLock.RLock()
	defer vw.wLock.RUnlock()

	var err error
	ws := memdb.NewWatchSet()

	vol, err = vw.state.CSIVolumeDenormalizePlugins(ws, vol.Copy())
	if err != nil {
		vw.logger.Error("could not query plugins for volume", "error", err)
		return nil
	}

	vol, err = vw.state.CSIVolumeDenormalize(ws, vol)
	if err != nil {
		vw.logger.Error("could not query allocs for volume", "error", err)
		return nil
	}
	return vol
}

// volumeReap collects errors for logging but doesn't return them
// to the main loop.
func (vw *volumeWatcher) volumeReap(vol *structs.CSIVolume) {
	err := vw.volumeReapImpl(vol)
	if err != nil {
		vw.logger.Error("error releasing volume claims", "error", err)
	}
}

func (vw *volumeWatcher) volumeReapImpl(vol *structs.CSIVolume) error {
	var result *multierror.Error
	nodeClaims := map[string]int{} // node IDs -> count
	collect := func(allocs map[string]*structs.Allocation,
		claims map[string]*structs.CSIVolumeClaim) {

		for allocID, alloc := range allocs {
			nodeClaims[alloc.NodeID]++
			if alloc.Terminated() {
				// don't overwrite the PastClaim if we've seen it before,
				// so that we can track state between subsequent calls
				_, exists := vol.PastClaims[allocID]
				if !exists {
					claim, ok := claims[allocID]
					if !ok {
						// COMPAT(1.0): the CSIVolumeClaim fields were added
						// after 0.11.1, so claims made before that may be
						// missing this value. note that we'll have non-nil
						// allocs here because we called denormalize on the
						// value.
						claim = &structs.CSIVolumeClaim{
							AllocationID: allocID,
							NodeID:       alloc.NodeID,
						}
					}
					claim.State = structs.CSIVolumeClaimStateTaken
					vol.PastClaims[claim.AllocationID] = claim
				}
			}
		}
	}

	collect(vol.ReadAllocs, vol.ReadClaims)
	collect(vol.WriteAllocs, vol.WriteClaims)

	if len(vol.PastClaims) == 0 {
		return nil
	}

	for _, claim := range vol.PastClaims {

		var err error

		// previous checkpoints may have set the past claim state already.
		// in practice we should never see CSIVolumeClaimStateControllerDetached
		// but having an option for the state makes it easy to add a checkpoint
		// in a backwards compatible way if we need one later
		switch claim.State {
		case structs.CSIVolumeClaimStateNodeDetached:
			goto NODE_DETACHED
		case structs.CSIVolumeClaimStateControllerDetached:
			goto RELEASE_CLAIM
		case structs.CSIVolumeClaimStateReadyToFree:
			goto RELEASE_CLAIM
		}

		err = vw.nodeDetach(vol, claim)
		if err != nil {
			result = multierror.Append(result, err)
			break
		}

	NODE_DETACHED:
		nodeClaims[claim.NodeID]--
		err = vw.controllerDetach(vol, claim, nodeClaims)
		if err != nil {
			result = multierror.Append(result, err)
			break
		}

	RELEASE_CLAIM:
		err = vw.checkpoint(vol, claim)
		if err != nil {
			result = multierror.Append(result, err)
			break
		}
		// the checkpoint deletes from the state store, but this operates
		// on our local copy which aids in testing
		delete(vol.PastClaims, claim.AllocationID)
	}

	return result.ErrorOrNil()

}

// nodeDetach makes the client NodePublish / NodeUnstage RPCs, which
// must be completed before controller operations or releasing the claim.
func (vw *volumeWatcher) nodeDetach(vol *structs.CSIVolume, claim *structs.CSIVolumeClaim) error {
	nReq := &cstructs.ClientCSINodeDetachVolumeRequest{
		PluginID:       vol.PluginID,
		VolumeID:       vol.ID,
		ExternalID:     vol.RemoteID(),
		AllocID:        claim.AllocationID,
		NodeID:         claim.NodeID,
		AttachmentMode: vol.AttachmentMode,
		AccessMode:     vol.AccessMode,
		ReadOnly:       claim.Mode == structs.CSIVolumeClaimRead,
	}

	err := vw.rpc.NodeDetachVolume(nReq,
		&cstructs.ClientCSINodeDetachVolumeResponse{})
	if err != nil {
		return fmt.Errorf("could not detach from node: %v", err)
	}
	claim.State = structs.CSIVolumeClaimStateNodeDetached
	return vw.checkpoint(vol, claim)
}

// controllerDetach makes the client RPC to the controller to
// unpublish the volume if a controller is required and no other
// allocs on the node need it
func (vw *volumeWatcher) controllerDetach(vol *structs.CSIVolume, claim *structs.CSIVolumeClaim, nodeClaims map[string]int) error {
	if !vol.ControllerRequired || nodeClaims[claim.NodeID] > 1 {
		claim.State = structs.CSIVolumeClaimStateReadyToFree
		return nil
	}

	// note: we need to get the CSI Node ID, which is not the same as
	// the Nomad Node ID
	ws := memdb.NewWatchSet()
	targetNode, err := vw.state.NodeByID(ws, claim.NodeID)
	if err != nil {
		return err
	}
	if targetNode == nil {
		return fmt.Errorf("%s: %s", structs.ErrUnknownNodePrefix, claim.NodeID)
	}
	targetCSIInfo, ok := targetNode.CSINodePlugins[vol.PluginID]
	if !ok {
		return fmt.Errorf("failed to find NodeInfo for node: %s", targetNode.ID)
	}

	plug, err := vw.state.CSIPluginByID(ws, vol.PluginID)
	if err != nil {
		return fmt.Errorf("plugin lookup error: %s %v", vol.PluginID, err)
	}
	if plug == nil {
		return fmt.Errorf("plugin lookup error: %s missing plugin", vol.PluginID)
	}

	cReq := &cstructs.ClientCSIControllerDetachVolumeRequest{
		VolumeID:        vol.RemoteID(),
		ClientCSINodeID: targetCSIInfo.NodeInfo.ID,
	}
	cReq.PluginID = plug.ID
	err = vw.rpc.ControllerDetachVolume(cReq,
		&cstructs.ClientCSIControllerDetachVolumeResponse{})
	if err != nil {
		return fmt.Errorf("could not detach from controller: %v", err)
	}
	claim.State = structs.CSIVolumeClaimStateReadyToFree
	return nil
}

func (vw *volumeWatcher) checkpoint(vol *structs.CSIVolume, claim *structs.CSIVolumeClaim) error {
	req := structs.CSIVolumeClaimRequest{
		VolumeID:     vol.ID,
		AllocationID: claim.AllocationID,
		NodeID:       claim.NodeID,
		Claim:        structs.CSIVolumeClaimRelease,
		State:        claim.State,
		WriteRequest: structs.WriteRequest{
			Namespace: vol.Namespace,
			// Region:    vol.Region, // TODO(tgross) should volumes have regions?
		},
	}
	index, err := vw.updateClaims([]structs.CSIVolumeClaimRequest{req})
	if err == nil && index != 0 {
		vw.wLock.Lock()
		defer vw.wLock.Unlock()
		vw.v.ModifyIndex = index
	}
	if err != nil {
		return fmt.Errorf("could not checkpoint claim release: %v", err)
	}
	return nil
}
