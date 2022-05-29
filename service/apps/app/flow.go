package app

import (
	"fmt"
	"strings"
)

// flow
// can be installed on all arch types
func (inst *Service) flow() {
	this.AppName = flow
	this.Owner = Owner
	this.Repo = flowRepo
	this.ServiceName = flowService
	this.RunAsUser = User
	this.Port = flowPort
	this.ServiceWorkingDirectory = fmt.Sprintf("%s/%s/%s", rootDir, appsInstallPath, this.AppName)
	this.AppsPath = fmt.Sprintf("%s/%s/%s", rootDir, appsInstallPath, this.AppName)
	this.AppPath = fmt.Sprintf("%s/%s", rootDir, this.AppName)
	this.ServiceExecStart = fmt.Sprintf("%s/app-amd64 -p %d -g %s -d data -prod", this.ServiceWorkingDirectory, this.Port, this.AppPath)
	if this.Arch == ArchAmd64 {
	} else {
		this.ServiceExecStart = strings.Replace(this.ServiceExecStart, "app-amd64", "app", 1)
	}
	return
}
