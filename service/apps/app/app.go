package app

import (
	"errors"
	"fmt"
	"gthub.com/NubeIO/rubix-cli-app/service/remote"
)

type App struct {
	Name                    string `json:"name"`     // rubix-wires
	AppName                 string `json:"app_name"` // rubixWires
	Owner                   string `json:"owner"`    // NubeIO
	Repo                    string `json:"repo"`     // wires-build
	Version                 string `json:"version"`
	ServiceName             string `json:"service_name"`     // nubeio-rubix-wires
	RunAsUser               string `json:"run_as_user"`      // root
	Port                    int    `json:"port"`             //1313
	MainDir                 string `json:"main_dir"`         // /data
	DataDir                 string `json:"data_dir"`         // /data
	ConfigDir               string `json:"config_dir"`       // MainDir+Repo+config
	ConfigFileName          string `json:"config_file_name"` // config.yml
	Description             string
	ServiceWorkingDirectory string `json:"service_working_directory"` // MainDir/apps/install/
	ServiceExecStart        string `json:"service_exec_start"`        // npm run prod:start --prod -- --datadir /data/rubix-wires/data --envFile /data/rubix-wires/config/.env
	HasConfig               bool   `json:"has_config"`
	ProductType             string `json:"product_type"`
	Arch                    string `json:"arch"`
}

/*
WIRES
service working dir
main |installDir                |repo         |version|name
/data/rubix-service/apps/install/wires-builds/v2.5.8/rubix-wires

exec start
-execDir | cmd                        | dataDir                        | config
/usr/bin/npm run prod:start --prod -- --datadir /data/rubix-wires/data --envFile /data/rubix-wires/config/.env

FLOW
service working dir
main | installDir               | repo         |version
/data/rubix-service/apps/install/flow-framework/v0.5.4


exec start
main |installDir                 | repo         |version|cmd|port     |main/repo           |data|production
/data/rubix-service/apps/install/flow-framework/v0.5.4/app -p 1660 -g /data/flow-framework -d data -prod

*/

const (
	Owner          = "NubeIO"
	MainDir        = "/data"
	MainInstallDir = "rubix-apps" // mainDir+/rubix-service
	User           = "root"
	ArchAmd64      = "amd64"
	ArchArm7       = "armv7"
)

//name
const (
	wires     = "rubix-wires"
	wiresPort = 1313
	flow      = "flow-framework"
	flowPort  = 1660
)

//repos
const (
	wiresRepo = "wires-builds"
	flowRepo  = flow
)

//ServiceName
const (
	wiresService = "nubeio-rubix-wires"
	flowService  = "nubeio-flow-framework"
)

const (
	configEnv = ".env"
	configYml = "config.yml"
)

func setDataDir(mainDir, name string) string {
	return fmt.Sprintf("%s/%s", mainDir, name)
}

func setConfigDir(mainDir, name string) string {
	return fmt.Sprintf("%s/%s/config", mainDir, name)
}

var this *App

func (inst *App) init() *App {
	this = &App{}
	return this
}

func (inst *App) GetApp() *App {
	return this
}

func (inst *App) GetExecStart() string {
	return this.ServiceExecStart
}

func (inst *App) initArchCheck() *remote.Admin {
	host := &remote.Admin{}
	return remote.New(host)
}

func (inst *App) setVersion(version string) *App {
	this.Version = version
	return this
}

func (inst *App) setArch(arch string) *App {
	this.Arch = arch
	return this
}

func (inst *App) getArch() (arch string, err error) {
	run := inst.initArchCheck()
	res, err := run.DetectModel()

	if res.IsArmv7l {
		arch = ArchArm7
	}
	if res.IsAMD64 {
		arch = ArchAmd64
	}
	return
}

func (inst *App) isLinux() (isLinux bool) {
	run := inst.initArchCheck()
	isLinux = run.ArchIsLinux()
	return
}

type Selection struct {
	AppName string
	Version string //version must be the installed version as in v0.0.1

}

func (inst *App) SelectApp(s *Selection) (*App, error) {
	inst.init()
	if !inst.isLinux() {
		return nil, errors.New("host mu by os type linux")
	}
	arch, err := inst.getArch()
	if err != nil {
		return nil, err
	}
	inst.setVersion(s.Version)
	inst.setArch(arch)
	switch s.AppName {
	case rubixWires.String():
		//wires can be installed on any linux arch
		inst.rubixWires()
		return inst.GetApp(), nil
	case flowFramework.String():
		//wires can be installed on any linux arch
		inst.flow()
		return inst.GetApp(), nil
	}
	return nil, errors.New("invalid app type, try rubixWires")

}
