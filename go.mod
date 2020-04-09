module github.com/hashicorp/nomad

go 1.14

require (
	cloud.google.com/go v0.44.4-0.20190827153918-f6872d26e209 // indirect
	contrib.go.opencensus.io/exporter/ocagent v0.4.12 // indirect
	github.com/Azure/azure-sdk-for-go v29.0.0+incompatible // indirect
	github.com/Azure/go-autorest v11.7.1+incompatible // indirect
	github.com/LK4D4/joincontext v0.0.0-20171026170139-1724345da6d5
	github.com/Microsoft/go-winio v0.4.15-0.20190919025122-fc70bd9a86b5
	github.com/Microsoft/hcsshim v0.8.8-0.20200312192636-fd0797d766b1 // indirect
	github.com/NVIDIA/gpu-monitoring-tools v0.0.0-20180829222009-86f2a9fac6c5
	github.com/NYTimes/gziphandler v1.1.1
	github.com/armon/circbuf v0.0.0-20150827004946-bbbad097214e
	github.com/armon/go-metrics v0.0.0-20190430140413-ec5e00d3c878
	github.com/aws/aws-sdk-go v1.25.41
	github.com/boltdb/bolt v1.3.1
	github.com/checkpoint-restore/go-criu v0.0.0-20190109184317-bdb7599cd87b // indirect
	github.com/container-storage-interface/spec v1.2.0-rc1.0.20191021210849-a33ece0a8a9f
	github.com/containerd/console v1.0.0 // indirect
	github.com/containerd/continuity v0.0.0-20200228182428-0f16d7a0959c // indirect
	github.com/containerd/go-cni v0.0.0-20190904155053-d20b7eebc7ee
	github.com/containernetworking/cni v0.7.2-0.20190612152420-dc953e2fd91f // indirect
	github.com/containernetworking/plugins v0.7.3-0.20190501191748-2d6d46d308b2
	github.com/coreos/go-iptables v0.4.3-0.20190724151750-969b135e941d
	github.com/coreos/go-semver v0.2.1-0.20170613092238-1817cd4bea52
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/cyphar/filepath-securejoin v0.2.3-0.20190205144030-7efe413b52e1 // indirect
	github.com/docker/cli v0.0.0-20200303215952-eb310fca4956
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v17.12.0-ce-rc1.0.20200330121334-7f8b4b621b5d+incompatible
	github.com/docker/docker-credential-helpers v0.6.2-0.20180719074751-73e5f5dbfea3 // indirect
	github.com/docker/go-units v0.4.0
	github.com/dustin/go-humanize v1.0.0
	github.com/elazarl/go-bindata-assetfs v1.0.0
	github.com/fatih/color v1.9.0
	github.com/fsouza/go-dockerclient v1.6.3
	github.com/godbus/dbus/v5 v5.0.3 // indirect
	github.com/golang/protobuf v1.3.4
	github.com/golang/snappy v0.0.1
	github.com/google/go-cmp v0.4.0
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/gorilla/websocket v1.4.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.1-0.20200228141219-3ce3d519df39
	github.com/hashicorp/consul v1.7.1-0.20200213195527-b137060630b4
	github.com/hashicorp/consul-template v0.24.1
	github.com/hashicorp/consul/api v1.4.0
	github.com/hashicorp/consul/sdk v0.4.0
	github.com/hashicorp/go-checkpoint v0.0.0-20171009173528-1545e56e46de
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/go-connlimit v0.2.0
	github.com/hashicorp/go-discover v0.0.0-20200108194735-7698de1390a1
	github.com/hashicorp/go-envparse v0.0.0-20180119215841-310ca1881b22
	github.com/hashicorp/go-getter v1.3.1-0.20190822194507-f5101da01173
	github.com/hashicorp/go-hclog v0.12.0
	github.com/hashicorp/go-immutable-radix v1.1.0
	github.com/hashicorp/go-memdb v1.0.3
	github.com/hashicorp/go-msgpack v1.1.5
	github.com/hashicorp/go-multierror v1.0.1-0.20191120192120-72917a1559e1
	github.com/hashicorp/go-plugin v1.0.2-0.20191004171845-809113480b55
	github.com/hashicorp/go-sockaddr v1.0.2
	github.com/hashicorp/go-syslog v1.0.0
	github.com/hashicorp/go-uuid v1.0.1
	github.com/hashicorp/go-version v1.2.1-0.20191009193637-2046c9d0f0b0
	github.com/hashicorp/golang-lru v0.5.3
	github.com/hashicorp/hcl v1.0.1-0.20191016231534-914dc3f8dd7c
	github.com/hashicorp/hcl2 v0.0.0-20190617160022-4fba5e1a75e3
	github.com/hashicorp/logutils v1.0.0
	github.com/hashicorp/memberlist v0.1.6
	github.com/hashicorp/net-rpc-msgpackrpc v0.0.0-20151116020338-a14192a58a69
	github.com/hashicorp/nomad/api v0.0.0-20200409153048-8259c6dfc051
	github.com/hashicorp/raft v1.1.3-0.20200211192230-365023de17e6
	github.com/hashicorp/raft-boltdb v0.0.0-20171010151810-6e5ba93211ea
	github.com/hashicorp/serf v0.8.5
	github.com/hashicorp/vault/api v1.0.5-0.20190730042357-746c0b111519
	github.com/hashicorp/vault/sdk v0.1.14-0.20190730042320-0dc007d98cc8
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d
	github.com/hpcloud/tail v1.0.1-0.20170814160653-37f427138745
	github.com/joyent/triton-go v0.0.0-20190112182421-51ffac552869 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pretty v0.2.0
	github.com/kr/pty v1.1.5
	github.com/kr/text v0.1.0
	github.com/mattn/go-colorable v0.1.4
	github.com/mitchellh/cli v1.0.0
	github.com/mitchellh/colorstring v0.0.0-20150917214807-8631ce90f286
	github.com/mitchellh/copystructure v1.0.0
	github.com/mitchellh/go-ps v0.0.0-20170309133038-4fdf99ab2936
	github.com/mitchellh/go-testing-interface v1.0.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/mrunalp/fileutils v0.0.0-20171103030105-7d4729fb3618 // indirect
	github.com/oklog/run v1.0.1-0.20180308005104-6934b124db28 // indirect
	github.com/onsi/ginkgo v1.10.1
	github.com/onsi/gomega v1.7.0
	github.com/opencontainers/runc v1.0.0-rc7
	github.com/opencontainers/runtime-spec v1.0.2-0.20200307132014-f49fed0d6290 // indirect
	github.com/opencontainers/selinux v1.4.1-0.20200311111634-a2f0d9c2aafc // indirect
	github.com/pkg/errors v0.9.1
	github.com/posener/complete v1.2.1
	github.com/prometheus/client_golang v0.9.4
	github.com/prometheus/common v0.4.1
	github.com/rs/cors v0.0.0-20170801073201-eabcc6af4bbe
	github.com/ryanuber/columnize v2.1.1-0.20170703205827-abc90934186a+incompatible
	github.com/ryanuber/go-glob v1.0.0
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529
	github.com/seccomp/libseccomp-golang v0.9.2-0.20200314001724-bdab42bd5128 // indirect
	github.com/shirou/gopsutil v2.20.2+incompatible
	github.com/sirupsen/logrus v1.4.3-0.20190518135202-2a22dbedbad1 // indirect
	github.com/skratchdot/open-golang v0.0.0-20160302144031-75fb7ed4208c
	github.com/stretchr/testify v1.5.1
	github.com/syndtr/gocapability v0.0.0-20170704070218-db04d3cc01c8
	github.com/vishvananda/netlink v1.1.0 // indirect
	github.com/zclconf/go-cty v1.0.0
	go.opencensus.io v0.22.1-0.20190713072201-b4a14686f0a9 // indirect
	golang.org/x/crypto v0.0.0-20200320181102-891825fb96df
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/api v0.9.1-0.20190824000815-035d22e00718 // indirect
	google.golang.org/genproto v0.0.0-20200302123026-7795fca6ccb1 // indirect
	google.golang.org/grpc v1.27.1
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
	gopkg.in/tomb.v2 v2.0.0-20140626144623-14b3d72120e8
	gotest.tools/v3 v3.0.2 // indirect
)

replace github.com/Microsoft/go-winio => github.com/endocrimes/go-winio v0.4.13-0.20190628114223-fb47a8b41948

replace github.com/hashicorp/nomad/api => ./api