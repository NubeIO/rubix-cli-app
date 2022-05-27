package apps

import (
	pprint "gthub.com/NubeIO/rubix-cli-app/pkg/helpers/print"
	"testing"
)

func TestLocalConnection(t *testing.T) {
	inst := &Installer{}
	timeout := 5
	service := "mosquitto"
	apps := New(inst)
	resp, _ := apps.ServiceStats(service, timeout)
	pprint.PrintJOSN(resp)
}
