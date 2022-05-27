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

func (inst *Apps) Start(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Start(service, systemOpts)
	if err != nil {
		resp.Err = err
		resp.Message = "tried to start the service but failed"
		return nil
	}
	resp.Ok = true
	resp.Message = "start service ok"
	return resp
}

func (inst *Apps) Stop(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Stop(service, systemOpts)
	if err != nil {
		resp.Err = err
		resp.Message = "tried to stop the service but failed"
		return nil
	}
	resp.Ok = true
	resp.Message = "stop service ok"
	return resp
}

func (inst *Apps) Status(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Message, err = systemctl.Status(service, systemOpts)
	if err != nil {
		resp.Err = err
		return resp
	}
	resp.Ok = true
	return resp
}

func (inst *Apps) Enable(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Enable(service, systemOpts)
	if err != nil {
		resp.Err = err
		resp.Message = "tried to enable the service but failed"
		return resp
	}
	resp.Ok = true
	resp.Message = "enabled the service ok"
	return resp
}

func (inst *Apps) Disable(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Disable(service, systemOpts)
	if err != nil {
		resp.Message = "disabled the service issue"
		resp.Err = err
		return resp
	}
	resp.Ok = true
	resp.Message = "disabled the service ok"
	return resp
}

func (inst *Apps) IsEnabled(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, err = systemctl.IsEnabled(service, systemOpts)
	if err != nil {
		resp.Message = "is not enabled"
		resp.Err = err
		return resp
	}
	resp.Message = "is enabled"
	return resp
}

func (inst *Apps) IsActive(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, resp.Message, err = systemctl.IsActive(service, systemOpts)
	if err != nil {
		resp.Err = err
		return resp
	}
	return resp
}

func (inst *Apps) IsRunning(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, resp.Message, err = systemctl.IsRunning(service, systemOpts)
	if err != nil {
		resp.Err = err
		return resp
	}
	return resp
}

func (inst *Apps) IsFailed(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, err = systemctl.IsFailed(service, systemOpts)
	if err != nil {
		resp.Message = "is failed"
		resp.Err = err
		return resp
	}
	resp.Message = "is ok"
	return resp
}

func (inst *Apps) IsInstalled(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, err = systemctl.IsInstalled(service, systemOpts)
	if err != nil {
		resp.Message = "is not installed"
		resp.Err = err
		return resp
	}
	resp.Message = "is installed"
	return resp
}

func (inst *Apps) ServiceStats(service string, timeout int) (resp systemctl.SystemState, err error) {
	systemOpts.Timeout = timeout
	resp, err = systemctl.State(service, systemOpts)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
