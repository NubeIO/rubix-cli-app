package apps

import (
	"github.com/NubeIO/lib-systemctl-go/systemctl"
)

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Err     error  `json:"err"`
}

var systemOpts = systemctl.Options{
	UserMode: false,
	Timeout:  0,
}

func (inst *Apps) Start(timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Start(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Err = err
		resp.Message = "tried to start  but failed"
		return resp
	}
	resp.Ok = true
	resp.Message = "start ok"
	return resp
}

func (inst *Apps) Stop(timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Stop(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Err = err
		resp.Message = "tried to stop but failed"
		return resp
	}
	resp.Ok = true
	resp.Message = "stop  ok"
	return resp
}

func (inst *Apps) Status(timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Message, err = systemctl.Status(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Err = err
		return resp
	}
	resp.Ok = true
	return resp
}

func (inst *Apps) Enable(timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Enable(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Err = err
		resp.Message = "tried to enable the but failed"
		return resp
	}
	resp.Ok = true
	resp.Message = "enabled ok"
	return resp
}

func (inst *Apps) Disable(timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Disable(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Message = "disabled issue"
		resp.Err = err
		return resp
	}
	resp.Ok = true
	resp.Message = "disabled ok"
	return resp
}

func (inst *Apps) IsEnabled(timeout int) (resp *Response, isEnabled bool) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, err = systemctl.IsEnabled(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Message = "is not enabled"
		resp.Err = err
		return resp, false
	}
	resp.Message = "is enabled"
	return resp, true
}

func (inst *Apps) IsActive(timeout int) (resp *Response, isActive bool) {

	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, resp.Message, err = systemctl.IsActive(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Err = err
		return resp, false
	}
	return resp, true
}

func (inst *Apps) IsRunning(timeout int) (resp *Response, isRunning bool) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, resp.Message, err = systemctl.IsRunning(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Err = err
		return resp, false
	}
	return resp, true
}

func (inst *Apps) IsFailed(timeout int) (resp *Response, isFailed bool) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, err = systemctl.IsFailed(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Message = "is failed"
		resp.Err = err
		return resp, true
	}
	resp.Message = "is ok"
	return resp, false
}

func (inst *Apps) IsInstalled(timeout int) (resp *Response, isInstall bool) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, err = systemctl.IsInstalled(inst.ServiceName, systemOpts)
	if err != nil {
		resp.Message = "is not installed"
		resp.Err = err
		return resp, false
	}
	resp.Message = "is installed"
	return resp, true
}

func (inst *Apps) ServiceNameStats(timeout int) (resp systemctl.SystemState, err error) {
	systemOpts.Timeout = timeout
	resp, err = systemctl.State(inst.ServiceName, systemOpts)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
