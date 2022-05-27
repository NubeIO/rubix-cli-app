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

func (inst *Installer) Start(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Start(service, systemOpts)
	if err != nil {
		resp.Err = err
		return nil
	}
	resp.Ok = true
	return resp
}

func (inst *Installer) Stop(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Stop(service, systemOpts)
	if err != nil {
		resp.Err = err
		return nil
	}
	resp.Ok = true
	return resp
}

func (inst *Installer) Status(service string, timeout int) (resp *Response) {
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

func (inst *Installer) Enable(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Enable(service, systemOpts)
	if err != nil {
		resp.Err = err
		return resp
	}
	resp.Ok = true
	return resp
}

func (inst *Installer) Disable(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	err := systemctl.Disable(service, systemOpts)
	if err != nil {
		resp.Err = err
		return resp
	}
	resp.Ok = true
	return resp
}

func (inst *Installer) IsEnabled(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, err = systemctl.IsEnabled(service, systemOpts)
	if err != nil {
		resp.Err = err
		return resp
	}
	resp.Ok = true
	return resp
}

func (inst *Installer) IsActive(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, resp.Message, err = systemctl.IsActive(service, systemOpts)
	if err != nil {
		resp.Err = err
		return resp
	}
	resp.Ok = true
	return resp
}

func (inst *Installer) IsFailed(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, err = systemctl.IsFailed(service, systemOpts)
	if err != nil {
		resp.Message = "is failed"
		resp.Err = err
		return resp
	}
	resp.Message = "is ok"
	resp.Ok = true
	return resp
}

func (inst *Installer) IsInstalled(service string, timeout int) (resp *Response) {
	resp = &Response{}
	systemOpts.Timeout = timeout
	resp.Ok, err = systemctl.IsInstalled(service, systemOpts)
	if err != nil {
		resp.Message = "is not installed"
		resp.Err = err
		return resp
	}
	resp.Message = "is installed"
	resp.Ok = true
	return resp
}

func (inst *Installer) ServiceStats(service string, timeout int) (resp systemctl.SystemStats, err error) {
	systemOpts.Timeout = timeout
	resp, err = systemctl.Stats(service, systemOpts)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
