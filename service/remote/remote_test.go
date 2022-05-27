package remote

import (
	"fmt"
	"gthub.com/NubeIO/rubix-cli-app/service/remote/old/ssh"
	"testing"

	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	"github.com/NubeIO/rubix-assist-model/model"
)

func TestLocalConnection(t *testing.T) {
	host := &Admin{
		SSH: &ssh.Host{
			Host: &model.Host{
				IsLocalhost: nils.NewBool(true),
			},
		},
	}
	run := New(host)
	run.Uptime()

}

func TestRemoteConnection(t *testing.T) {
	host := &Admin{
		SSH: &ssh.Host{
			Host: &model.Host{
				IP:       "192.168.15.103",
				Port:     22,
				Username: "debian",
				Password: "N00B2828",
			},
		},
	}
	run := New(host)
	out, err := run.EdgeSetIP(&EdgeNetworking{IPAddress: "192.168.15.103", SubnetMask: "255.255.255.0", Gateway: "192.168.15.1"})
	fmt.Println(out, err)

}
