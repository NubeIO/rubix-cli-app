package app

import "fmt"

func (inst *Service) rubixWires() {
	this.Name = wires
	this.AppName = RubixWires.String()
	this.Owner = Owner
	this.Repo = wiresRepo
	this.ServiceName = wiresService
	this.RunAsUser = User
	this.Port = wiresPort
	this.DataDir = fmt.Sprintf("%s/%s/data", rootDir, wires)
	this.ConfigDir = fmt.Sprintf("%s/%s/config", rootDir, wires)
	this.ConfigFileName = configEnv
	this.ServiceWorkingDirectory = fmt.Sprintf("%s/%s/%s/%s/%s", rootDir, mainInstallDir, wiresRepo, this.Version, wires)
	this.ServiceExecStart = fmt.Sprintf("/usr/bin/npm run prod:start --prod --datadir %s --envFile %s/%s", this.DataDir, this.ConfigDir, configEnv)
	this.HasConfig = true
	return
}
