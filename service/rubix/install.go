package rubix

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-edge/pkg/logger"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
	"time"
)

type Install struct {
	Name      string `json:"name"`
	BuildName string `json:"build_name"`
	Version   string `json:"version"`
	Source    string `json:"source"`
}

type Response struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (inst *App) InstallApp(app *Install) (*Response, error) {
	var appName = app.Name
	var appBuildName = app.BuildName
	var version = app.Version
	var source = app.Source
	if err := inst.installApp(appName, appBuildName, version, source); err != nil {
		return nil, err
	}
	return &Response{
		Name:    appName,
		Message: "installed app ok",
	}, nil

}

// InstallApp make all the required dirs and unzip build
//	zip, pass in the zip folder, or you can pass in a local path to param localZip
func (inst *App) installApp(appName, appBuildName, version string, source string) error {
	// make the dirs
	err := inst.DirsInstallApp(appName, appBuildName, version)
	if err != nil {
		return err
	}
	log.Infof("made all dirs for app:%s,  buildName:%s, version:%s", appName, appBuildName, version)
	dest := inst.getAppInstallPathAndVersion(appBuildName, version)
	log.Infof("app zip source:%s", source)
	log.Infof("app zip dest:%s", dest)
	// unzip the build to the app dir  /data/rubix-service/install/wires-build
	_, err = inst.unZip(source, dest)
	if err != nil {
		return err
	}
	return nil
}

func (inst *App) InstallService(app App, serviceFile *multipart.FileHeader, localServiceFile string) (*ctl.InstallResp, error) {
	var err error
	//found := inst.ConfirmAppDir(app.Name)
	//found = inst.ConfirmAppInstallDir(app.AppBuildName)
	var serviceName = app.ServiceName
	var fileSource string
	if localServiceFile != "" { //
		source := fmt.Sprintf("%s/%s", HostDownloadPath, localServiceFile)
		dest := fmt.Sprintf("%s/%s", TmpDir, localServiceFile)
		fileSource = dest
		err := moveFile(source, dest, false)
		if err != nil {
			log.Errorf("move zip:%s: err:%s", dest, err.Error())
			//return err
		}
	} else {
		// save app in tmp dir
		fileSource, err = inst.saveUploadedFile(serviceFile, TmpDir)
		if err != nil {
			log.Errorf("move zip:%s: err:%s", fileSource, err.Error())
			//return err
		}
	}

	service, err := inst.installService(serviceName, fileSource)
	if err != nil {
		return nil, err
	}

	return service, err
}

// InstallService a new linux service
//	- service: the service name (eg: nubeio-rubix-wires)
//	- path: the service file path and name (eg: "/tmp/rubix-bios.service")
func (inst *App) installService(service, tmpServiceFile string) (*ctl.InstallResp, error) {
	var err error
	resp := &ctl.InstallResp{}
	// path := "/tmp/nubeio-rubix-bios.service"
	timeOut := 30
	ser := ctl.New(service, tmpServiceFile)
	ser.InstallOpts = ctl.InstallOpts{
		Options: systemctl.Options{Timeout: timeOut},
	}
	err = ser.TransferFile()
	if err != nil {
		fmt.Println("full install error", err)
		return nil, err
	}

	resp = ser.Install()
	if err != nil {
		fmt.Println("full install error", err)
		return nil, err
	}
	time.Sleep(8 * time.Second)
	active, status, err := systemctl.IsRunning(service, systemctl.Options{})
	if err != nil {
		logger.Logger.Errorf("service found or failed to check IsRunning: %s: %v", service, err)
		return nil, err
	}
	logger.Logger.Infof("service: %s: isActive: %t status: %s", service, active, status)
	return resp, nil
}
