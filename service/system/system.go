package system

import (
	"errors"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
)

type Apps struct {
	Token       string `json:"token"`   // git token
	Version     string `json:"version"` // version to install
	Perm        int
	ServiceName string
}

var err error

func New(inst *Apps) (*Apps, error) {
	if inst == nil {
		return nil, errors.New("type apps must not be nil")
	}
	if inst.Perm == 0 {
		inst.Perm = Permission
	}
	return inst, err
}

const Permission = 0700

func (inst *Apps) SystemCtlAction(action string, timeout int) (*SystemResponse, error) {
	actionResp := &SystemResponse{}
	switch action {
	case start.String():
		return inst.Start(timeout)
	case stop.String():
		return inst.Stop(timeout)
	case enable.String():
		return inst.Enable(timeout)
	case disable.String():
		return inst.Disable(timeout)
	}
	return actionResp, errors.New("no valid action found try, start, stop, enable or disable")
}

func (inst *Apps) SystemCtlStatus(action string, timeout int) (*SystemResponseChecks, error) {
	actionResp := &SystemResponseChecks{}
	switch action {
	case isRunning.String():
		return inst.IsRunning(timeout)
	case isInstalled.String():
		return inst.IsInstalled(timeout)
	case isEnabled.String():
		return inst.IsEnabled(timeout)
	case isActive.String():
		return inst.IsActive(timeout)
	case isFailed.String():
		return inst.IsFailed(timeout)
	}
	return actionResp, errors.New("no valid action found try, isRunning, isInstalled, isEnabled, isActive or isFailed")
}

type Mass struct {
	Apps    []string
	Action  string
	Timeout int `json:"timeout"`
}

type massResponse struct {
	AppName string `json:"app_name"`
	Action  string `json:"action"`
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Err     error  `json:"err"`
}

func (inst *Apps) Mass(mass *Mass) ([]massResponse, error) {
	var response []massResponse
	// for _, app := range mass.Apps {
	//	appService := &AppService{}
	//	actionType := appService.Action
	//	actionResp, err := inst.Action(appService)
	//	if err != nil {
	//		return nil, err
	//	}
	//	res := massResponse{
	//		AppName: app,
	//		Action:  actionType,
	//		Ok:      actionResp.Ok,
	//		Message: actionResp.Message,
	//		Err:     actionResp.Err,
	//	}
	//	httpresp = append(httpresp, res)
	// }
	return response, nil
}

type SystemResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

var DefaultTimeout = 30

var systemOpts = systemctl.Options{
	UserMode: false,
	Timeout:  DefaultTimeout,
}

func (inst *Apps) Start(timeout int) (resp *SystemResponse, err error) {
	resp = &SystemResponse{}
	systemOpts.Timeout = timeout
	err = systemctl.Start(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Message = "tried to start but failed"
		return resp, err
	}
	resp.Ok = true
	resp.Message = "start ok"
	return resp, nil
}

func (inst *Apps) Stop(timeout int) (resp *SystemResponse, err error) {
	resp = &SystemResponse{}
	systemOpts.Timeout = timeout
	err = systemctl.Stop(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Message = "tried to stop but failed"
		return resp, err
	}
	resp.Ok = true
	resp.Message = "stop ok"
	return resp, nil
}

func (inst *Apps) Enable(timeout int) (resp *SystemResponse, err error) {
	resp = &SystemResponse{}
	systemOpts.Timeout = timeout
	err = systemctl.Enable(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Message = "tried to enable but failed"
		return resp, err
	}
	resp.Ok = true
	resp.Message = "enable ok"
	return resp, nil
}

func (inst *Apps) Disable(timeout int) (resp *SystemResponse, err error) {
	resp = &SystemResponse{}
	systemOpts.Timeout = timeout
	err = systemctl.Disable(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Message = "tried to disable but failed"
		return resp, err
	}
	resp.Ok = true
	resp.Message = "disable ok"
	return resp, nil
}

type SystemResponseChecks struct {
	Is      bool   `json:"is"`
	Message string `json:"message"`
}

func (inst *Apps) IsEnabled(timeout int) (resp *SystemResponseChecks, err error) {
	resp = &SystemResponseChecks{}
	systemOpts.Timeout = timeout
	out, err := systemctl.IsEnabled(inst.ServiceName, systemOpts)
	if err != nil || out == false {
		resp.Message = "is not enabled"
		return resp, err
	}
	resp.Is = out
	resp.Message = "is enabled"
	return
}

func (inst *Apps) IsFailed(timeout int) (resp *SystemResponseChecks, err error) {
	resp = &SystemResponseChecks{}
	systemOpts.Timeout = timeout
	out, err := systemctl.IsFailed(inst.ServiceName, systemOpts)
	if err != nil || out == true {
		resp.Message = "is failed"
		return resp, err
	}
	resp.Is = out
	resp.Message = "is not failed"
	return
}

func (inst *Apps) IsInstalled(timeout int) (resp *SystemResponseChecks, err error) {
	resp = &SystemResponseChecks{}
	systemOpts.Timeout = timeout
	out, err := systemctl.IsInstalled(inst.ServiceName, systemOpts)
	if err != nil || out == false {
		resp.Message = "is not installed"
		return resp, err
	}
	resp.Is = out
	resp.Message = "is installed"
	return
}

func (inst *Apps) IsActive(timeout int) (resp *SystemResponseChecks, err error) {
	resp = &SystemResponseChecks{}
	systemOpts.Timeout = timeout
	out, msg, err := systemctl.IsActive(inst.ServiceName, systemOpts)
	if err != nil || out == false {
		resp.Message = msg
		return resp, err
	}
	resp.Is = out
	resp.Message = msg
	return
}

func (inst *Apps) IsRunning(timeout int) (resp *SystemResponseChecks, err error) {
	resp = &SystemResponseChecks{}
	systemOpts.Timeout = timeout
	out, msg, err := systemctl.IsRunning(inst.ServiceName, systemOpts)
	if err != nil || out == false {
		resp.Message = msg
		return resp, err
	}
	resp.Is = out
	resp.Message = msg
	return
}

func (inst *Apps) Status(timeout int) (message string, err error) {
	systemOpts.Timeout = timeout
	return systemctl.Status(inst.ServiceName, systemOpts)
}

func (inst *Apps) ServiceStats(timeout int) (resp systemctl.SystemState, err error) {
	systemOpts.Timeout = timeout
	resp, err = systemctl.State(inst.ServiceName, systemOpts)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
