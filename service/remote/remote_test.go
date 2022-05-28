package remote

import (
	pprint "gthub.com/NubeIO/rubix-cli-app/pkg/helpers/print"
	"testing"
)

func TestLocalConnection(t *testing.T) {

}

func TestRemoteConnection(t *testing.T) {
	host := &Admin{}
	run := New(host)
	run.ArchIsLinux()

	r, rr := run.DetectArch()

	pprint.PrintJOSN(r)
	pprint.PrintJOSN(rr)
	//out, err := run.EdgeSetIP(&EdgeNetworking{IPAddress: "192.168.15.103", SubnetMask: "255.255.255.0", Gateway: "192.168.15.1"})
	//fmt.Println(out, err)

}
