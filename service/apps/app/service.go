package app

import (
	"errors"
	"fmt"
	"gthub.com/NubeIO/rubix-cli-app/service/remote"
)

type Service struct {
	AppName                 string `json:"app_name"` // rubix-wires
	Owner                   string `json:"owner"`    // NubeIO
	Repo                    string `json:"repo"`     // wires-build
	Version                 string `json:"version"`
	ServiceName             string `json:"service_name"`     // nubeio-rubix-wires
	RunAsUser               string `json:"run_as_user"`      // root
	Port                    int    `json:"port"`             //1313
	AppsPath                string `json:"apps_path"`        // /data/rubix-apps/install/flow-framework
	AppPath                 string `json:"app_path"`         // /data/flow-framework
	DataDir                 string `json:"data_dir"`         // /data
	ConfigDir               string `json:"config_dir"`       // MainDir+Repo+config
	ConfigFileName          string `json:"config_file_name"` // config.yml
	ServiceDescription      string `json:"description"`
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

var (
	rootDir         = "/data"
	appsPath        = "rubix-apps" // mainDir+/rubix-service
	appsInstallPath = fmt.Sprintf("%s/installed", appsPath)
)

const (
	Owner     = "NubeIO"
	User      = "root"
	ArchAmd64 = "amd64"
	ArchArm7  = "armv7"
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

func (inst *Service) init() *Service {
	return this
}

func (inst *Service) initArchCheck() *remote.Admin {
	host := &remote.Admin{}
	return remote.New(host)
}

func (inst *Service) setVersion(version string) *Service {
	this.Version = version
	return this
}

func (inst *Service) setArch(arch string) *Service {
	this.Arch = arch
	return this
}

func (inst *Service) getArch() (arch string, err error) {
	run := inst.initArchCheck()
	res, err := run.DetectArch()

	if res.IsArmv7l {
		arch = ArchArm7
	}
	if res.IsAMD64 {
		arch = ArchAmd64
	}
	return
}

func (inst *Service) isLinux() (isLinux bool) {
	run := inst.initArchCheck()
	isLinux = run.ArchIsLinux()
	return
}

func (inst *Service) SelectApp() (*Service, error) {
	name, _, err := CheckAppName(appName)
	if err != nil {
		return nil, err
	}
	inst.init()
	if !inst.isLinux() {
		return nil, errors.New("host must by an os type linux")
	}
	arch, err := inst.getArch()
	if err != nil {
		return nil, err
	}
	inst.setVersion(appVersion)
	inst.setArch(arch)
	switch name {
	case RubixWires:
		//wires can be installed on any linux arch
		inst.rubixWires()
		if this.ServiceName == "" {
			return nil, errors.New("service name can not be nil, try nubeio-rubix-wires")
		}
		return this, nil
	case FlowFramework:
		//wires can be installed on any linux arch
		inst.flow()
		if this.ServiceName == "" {
			return nil, errors.New("service name can not be nil, try nubeio-rubix-wires")
		}
		return this, nil
	}
	return nil, errors.New("invalid app type, try rubixWires")

}
