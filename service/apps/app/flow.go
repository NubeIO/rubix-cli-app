package app

import "fmt"

func (inst *App) flow() {
	this.Name = flow
	this.AppName = flowFramework.String()
	this.Owner = Owner
	this.Repo = flowRepo
	this.ServiceName = flowService
	this.RunAsUser = User
	this.Port = flowPort
	this.MainDir = flow
	this.DataDir = fmt.Sprintf("%s/%s/data", this.MainDir, wires)
	this.ConfigDir = fmt.Sprintf("%s/%s/config", this.MainDir, wires)
	this.ConfigFileName = configYml
	this.ServiceWorkingDirectory = fmt.Sprintf("%s/%s/%s/%s/%s", this.MainDir, MainInstallDir, wiresRepo, this.Version, wires)
	this.ServiceExecStart = fmt.Sprintf("/usr/bin/npm run prod:start --prod --datadir %s --envFile %s/%s", this.DataDir, this.ConfigDir, configEnv)
	this.HasConfig = true
	return
}
