package apps

import (
	"github.com/NubeIO/lib-systemctl-go/systemctl"
)

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

var defaultTimeout = 30

var systemOpts = systemctl.Options{
	UserMode: false,
	Timeout:  defaultTimeout,
}

func (inst *Apps) Start(timeout int) (resp *Response, err error) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err = systemctl.Start(inst.serviceName, systemOpts)
	if err != nil {
		resp.Message = "tried to start but failed"
		return resp, err
	}
	resp.Ok = true
	resp.Message = "start ok"
	return resp, nil
}

func (inst *Apps) Stop(timeout int) (resp *Response, err error) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err = systemctl.Stop(inst.serviceName, systemOpts)
	if err != nil {
		resp.Message = "tried to stop but failed"
		return resp, err
	}
	resp.Ok = true
	resp.Message = "stop ok"
	return resp, nil
}

func (inst *Apps) Enable(timeout int) (resp *Response, err error) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err = systemctl.Enable(inst.serviceName, systemOpts)
	if err != nil {
		resp.Message = "tried to enable but failed"
		return resp, err
	}
	resp.Ok = true
	resp.Message = "enable ok"
	return resp, nil
}

func (inst *Apps) Disable(timeout int) (resp *Response, err error) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err = systemctl.Disable(inst.serviceName, systemOpts)
	if err != nil {
		resp.Message = "tried to disable but failed"
		return resp, err
	}
	resp.Ok = true
	resp.Message = "disable ok"
	return resp, nil
}

func (inst *Apps) IsEnabled(timeout int) (out bool, err error) {
	systemOpts.Timeout = timeout
	return systemctl.IsEnabled(inst.serviceName, systemOpts)
}

func (inst *Apps) IsFailed(timeout int) (out bool, err error) {
	systemOpts.Timeout = timeout
	return systemctl.IsFailed(inst.serviceName, systemOpts)
}

func (inst *Apps) IsInstalled(timeout int) (out bool, err error) {
	systemOpts.Timeout = timeout
	return systemctl.IsInstalled(inst.serviceName, systemOpts)
}

func (inst *Apps) IsActive(timeout int) (active bool, status string, err error) {
	systemOpts.Timeout = timeout
	return systemctl.IsActive(inst.serviceName, systemOpts)
}

func (inst *Apps) IsRunning(timeout int) (active bool, status string, err error) {
	systemOpts.Timeout = timeout
	return systemctl.IsRunning(inst.serviceName, systemOpts)
}

func (inst *Apps) Status(timeout int) (message string, err error) {
	systemOpts.Timeout = timeout
	return systemctl.Status(inst.serviceName, systemOpts)
}

func (inst *Apps) serviceNameStats(timeout int) (resp systemctl.SystemState, err error) {
	systemOpts.Timeout = timeout
	resp, err = systemctl.State(inst.serviceName, systemOpts)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
