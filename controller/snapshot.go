package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-edge/model"
	"github.com/NubeIO/rubix-edge/pkg/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var systemPath = "/lib/systemd/system"
var dataFolder = "data"
var systemFolder = "system"
var createStatus = "N/A"
var restoreStatus = "N/A"

func (inst *Controller) Create(c *gin.Context) {
	if createStatus == "Running" {
		responseHandler(nil, errors.New("create snapshot is running"), c)
		return
	}
	createStatus = "Running"
	deviceInfo, err := inst.RubixRegistry.GetDeviceInfo()
	if err != nil {
		createStatus = "Failed"
		responseHandler(nil, err, c)
		return
	}
	destinationPath := fmt.Sprintf("%s/%s-%s-%s_%s", config.Config.GetAbsTempDir(), deviceInfo.ClientName,
		deviceInfo.SiteName, deviceInfo.DeviceName, time.Now().UTC().Format("20060102T150405"))
	_ = copyFolder(config.Config.GetAbsDataDir(), path.Join(destinationPath, dataFolder))

	systemFiles, err := filepath.Glob(path.Join(systemPath, "nubeio-*"))
	if err != nil {
		createStatus = "Failed"
		responseHandler(nil, err, c)
		return
	}
	copyFiles(systemFiles, path.Join(destinationPath, systemFolder))

	zipDestinationPath := destinationPath + ".zip"
	err = fileutils.RecursiveZip(destinationPath, zipDestinationPath)
	if err != nil {
		createStatus = "Failed"
		responseHandler(nil, err, c)
		return
	}
	err = checkSnapshotSize(zipDestinationPath)
	if err != nil {
		createStatus = "Failed"
		responseHandler(nil, err, c)
		return
	}
	createStatus = "Created"
	c.FileAttachment(zipDestinationPath, filepath.Base(zipDestinationPath))
}

func (inst *Controller) Restore(c *gin.Context) {
	if restoreStatus == "Running" {
		responseHandler(nil, errors.New("restore snapshot is running"), c)
		return
	}
	restoreStatus = "Running"
	useGlobalUUID, _ := toBool(c.Query("use_global_uuid"))
	file, _ := c.FormFile("file")
	destinationFilePath := path.Join(config.Config.GetAbsTempDir(), file.Filename)
	err := c.SaveUploadedFile(file, destinationFilePath)
	if err != nil {
		restoreStatus = err.Error()
		responseHandler(nil, err, c)
		return
	}
	_, err = fileutils.Unzip(destinationFilePath, path.Join(config.Config.GetAbsTempDir(), ""), os.FileMode(inst.FileMode))
	if err != nil {
		restoreStatus = err.Error()
		responseHandler(nil, err, c)
		return
	}
	unzippedFolderPath := path.Join(config.Config.GetAbsTempDir(), fileNameWithoutExtension(file.Filename))
	services, _ := fileutils.ListFiles(path.Join(unzippedFolderPath, systemFolder))
	inst.stopServices(services)
	err = copyFolder(path.Join(unzippedFolderPath, systemFolder), systemPath)
	if err != nil {
		restoreStatus = err.Error()
		responseHandler(nil, err, c)
		return
	}
	err = copyFolder(path.Join(unzippedFolderPath, dataFolder), config.Config.GetAbsDataDir())
	if err != nil {
		restoreStatus = err.Error()
		responseHandler(nil, err, c)
		return
	}
	if !useGlobalUUID {
		inst.updateDeviceInfo()
	}
	inst.enableAndRestartServices(services)
	message := model.Message{Message: "snapshot restored successfully"}
	restoreStatus = "Restored"
	responseHandler(message, err, c)
}

func (inst *Controller) Status(c *gin.Context) {
	responseHandler(model.SnapshotStatus{Create: createStatus, Restore: restoreStatus}, nil, c)
}

func (inst *Controller) stopServices(services []string) {
	var wg sync.WaitGroup
	for _, service := range services {
		wg.Add(1)
		service := service
		go func() {
			defer wg.Done()
			if !strings.Contains(service, "rubix-edge") {
				err := inst.SystemCtl.Stop(service)
				if err != nil {
					log.Errorf("err: %s", err.Error())
				}
			}
		}()
	}
	wg.Wait()
}

func (inst *Controller) enableAndRestartServices(services []string) {
	var wg sync.WaitGroup
	for _, service := range services {
		wg.Add(1)
		service := service
		go func() {
			defer wg.Done()
			if !strings.Contains(service, "rubix-edge") {
				err := inst.SystemCtl.Enable(service)
				if err != nil {
					log.Errorf("err: %s", err.Error())
				}
				err = inst.SystemCtl.Restart(service)
				if err != nil {
					log.Errorf("err: %s", err.Error())
				}
			}
		}()
	}
	wg.Wait()
}

func (inst *Controller) updateDeviceInfo() {
	deviceInfo, err := inst.RubixRegistry.GetDeviceInfo()
	if err != nil {
		log.Errorf("err: %s", err.Error())
	} else {
		deviceInfo.GlobalUUID = ""
		deviceInfoDefaultRaw, err := json.Marshal(deviceInfo)
		if err != nil {
			log.Errorf("err: %s", err.Error())
		}
		err = os.WriteFile(path.Join("/", config.Config.GetAbsDataDir(), "rubix-registry/device_info.json"),
			deviceInfoDefaultRaw, os.FileMode(inst.FileMode))
		if err != nil {
			log.Errorf("err: %s", err.Error())
		}
	}
}
