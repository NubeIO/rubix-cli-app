package apps

import (
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
)

/*
- stop, disable service
- remove service file
- delete files from data as in db, /data/flow-framework (optional)
*/

// UninstallService
//	- service nubeio-flow-framework
func (inst *Apps) UninstallService(service string) (*ctl.RemoveRes, error) {
	resp := &ctl.RemoveRes{}
	ser := ctl.New(service, "")
	ser.InstallOpts = ctl.InstallOpts{
		Options: systemctl.Options{Timeout: DefaultTimeout},
	}
	resp, err := ser.Remove()
	return resp, err
}
