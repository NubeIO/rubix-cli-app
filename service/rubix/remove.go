package rubix

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
)

type RemoveRes struct {
	DeleteAppDir         string `json:"delete_app_dir"`
	DeleteAppInstallDir  string `json:"delete_app_install_dir"`
	ServiceWasInstalled  string `json:"service_was_installed"`
	Stop                 bool   `json:"stop"`
	Disable              bool   `json:"disable"`
	DaemonReload         bool   `json:"daemon_reload"`
	RestartFailed        bool   `json:"restart_failed"`
	DeleteServiceFile    bool   `json:"delete_service_file"`
	DeleteServiceFileUsr bool   `json:"delete_service_file_usr"`
	Error                string `json:"error,omitempty"`
}

/*
- stop, disable service
- remove service file
*/

// UninstallService
//	- service nubeio-flow-framework
func (inst *App) UninstallService(appName, appBuildName, service string) (*RemoveRes, error) {
	ser := ctl.New(service, "")
	ser.InstallOpts = ctl.InstallOpts{
		Options: systemctl.Options{Timeout: DefaultTimeout},
	}
	remove, _ := ser.Remove()
	resp := &RemoveRes{
		ServiceWasInstalled:  remove.ServiceWasInstalled,
		Stop:                 remove.Stop,
		Disable:              remove.Disable,
		DaemonReload:         remove.DaemonReload,
		RestartFailed:        remove.RestartFailed,
		DeleteServiceFile:    remove.DeleteServiceFile,
		DeleteServiceFileUsr: remove.DeleteServiceFileUsr,
	}
	err := inst.RemoveApp(appName)
	var removeApp = "removed app from data dir ok"
	var removeAppInstall = "removed app from install dir ok"
	if err != nil {
		resp.Error = err.Error()
		removeApp = fmt.Sprintf("failed to delete app from data dir")
	}
	err = inst.RemoveAppInstall(appBuildName)
	if err != nil {
		resp.Error = err.Error()
		removeAppInstall = fmt.Sprintf("failed to delete app from install dir")
	}
	resp.DeleteAppDir = removeApp
	resp.DeleteAppInstallDir = removeAppInstall
	return resp, nil
}

// RemoveApp delete app
func (inst *App) RemoveApp(appName string) error {
	return inst.RmRF(inst.getAppPath(appName))
}

// RemoveAppInstall delete app install path
func (inst *App) RemoveAppInstall(appBuildName string) error {
	return inst.RmRF(inst.getAppInstallPath(appBuildName))
}

// RmRF remove file and all files inside
func (inst *App) RmRF(path string) error {
	return fileutils.New().RmRF(path)

}