package app

import "fmt"

func (inst *App) rubixWires() {
	this.Name = wires
	this.AppName = rubixWires.String()
	this.Owner = Owner
	this.Repo = wiresRepo
	this.ServiceName = wiresService
	this.RunAsUser = User
	this.Port = wiresPort
	this.MainDir = wires
	this.DataDir = fmt.Sprintf("%s/%s/data", this.MainDir, wires)
	this.ConfigDir = fmt.Sprintf("%s/%s/config", this.MainDir, wires)
	this.ConfigFileName = configEnv
	this.ServiceWorkingDirectory = fmt.Sprintf("%s/%s/%s/%s/%s", this.MainDir, MainInstallDir, wiresRepo, this.Version, wires)
	this.ServiceExecStart = fmt.Sprintf("/usr/bin/npm run prod:start --prod --datadir %s --envFile %s/%s", this.DataDir, this.ConfigDir, configEnv)
	this.HasConfig = true
	return
}
