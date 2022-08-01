package apps

import "github.com/NubeIO/lib-systemctl-go/systemctl"

type CtlBody struct {
	Service      string   `json:"service"`
	Action       string   `json:"action"`
	Timeout      int      `json:"timeout"`
	ServiceNames []string `json:"service_names"`
}

func (inst *EdgeApps) CtlAction(body *CtlBody) (*systemctl.SystemResponse, error) {
	return inst.Ctl.CtlAction(body.Action, body.Service, body.Timeout)
}

func (inst *EdgeApps) CtlStatus(body *CtlBody) (*systemctl.SystemResponseChecks, error) {
	return inst.Ctl.CtlStatus(body.Action, body.Service, body.Timeout)
}

func (inst *EdgeApps) ServiceMassAction(body *CtlBody) ([]systemctl.MassSystemResponse, error) {
	return inst.Ctl.ServiceMassAction(body.ServiceNames, body.Action, body.Timeout)
}

func (inst *EdgeApps) ServiceMassCheck(body *CtlBody) ([]systemctl.MassSystemResponseChecks, error) {
	return inst.Ctl.ServiceMassCheck(body.ServiceNames, body.Action, body.Timeout)
}
