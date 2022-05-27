package remote

import (
	"fmt"
	"testing"
)

func TestLocalConnection(t *testing.T) {

}

func TestRemoteConnection(t *testing.T) {
	host := &Admin{}
	run := New(host)
	out, err := run.EdgeSetIP(&EdgeNetworking{IPAddress: "192.168.15.103", SubnetMask: "255.255.255.0", Gateway: "192.168.15.1"})
	fmt.Println(out, err)

}
