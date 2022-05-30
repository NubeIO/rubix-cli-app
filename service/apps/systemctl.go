package apps

import (
	"github.com/NubeIO/lib-systemctl-go/systemctl"
)

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
	err = systemctl.Start(inst.App.ServiceName, systemOpts)
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
	err = systemctl.Stop(inst.App.ServiceName, systemOpts)
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
	err = systemctl.Enable(inst.App.ServiceName, systemOpts)
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
	err = systemctl.Disable(inst.App.ServiceName, systemOpts)
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
	out, err := systemctl.IsEnabled(inst.App.ServiceName, systemOpts)
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
	out, err := systemctl.IsFailed(inst.App.ServiceName, systemOpts)
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
	out, err := systemctl.IsInstalled(inst.App.ServiceName, systemOpts)
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
	out, msg, err := systemctl.IsActive(inst.App.ServiceName, systemOpts)
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
	out, msg, err := systemctl.IsRunning(inst.App.ServiceName, systemOpts)
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
	return systemctl.Status(inst.App.ServiceName, systemOpts)
}

func (inst *Apps) ServiceStats(timeout int) (resp systemctl.SystemState, err error) {
	systemOpts.Timeout = timeout
	resp, err = systemctl.State(inst.App.ServiceName, systemOpts)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
