package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-edge/model"
	"github.com/NubeIO/rubix-edge/pkg/config"
	"github.com/NubeIO/rubix-edge/utils"
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

var createStatus = model.CreateNotAvailable
var restoreStatus = model.RestoreNotAvailable

func (inst *Controller) CreateSnapshot(c *gin.Context) {
	if createStatus == model.Creating {
		responseHandler(nil, errors.New("creating snapshot"), c)
		return
	}
	createStatus = model.Creating
	deviceInfo, err := inst.RubixRegistry.GetDeviceInfo()
	if err != nil {
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	filePrefix := fmt.Sprintf("%s-%s-%s", deviceInfo.ClientName, deviceInfo.SiteName, deviceInfo.DeviceName)
	previousFiles, _ := filepath.Glob(path.Join(config.Config.GetAbsTempDir(), fmt.Sprintf("%s*", filePrefix)))
	utils.DeleteFiles(previousFiles, config.Config.GetAbsTempDir())

	destinationPath := fmt.Sprintf("%s/%s_%s", config.Config.GetAbsTempDir(), filePrefix,
		time.Now().UTC().Format("20060102T150405"))
	_ = utils.CopyDir(config.Config.GetAbsDataDir(), path.Join(destinationPath, dataFolder))

	systemFiles, err := filepath.Glob(path.Join(systemPath, "nubeio-*"))
	if err != nil {
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	utils.CopyFiles(systemFiles, path.Join(destinationPath, systemFolder))

	zipDestinationPath := destinationPath + ".zip"
	err = fileutils.RecursiveZip(destinationPath, zipDestinationPath)
	if err != nil {
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	_ = os.RemoveAll(destinationPath)
	createStatus = model.Created
	c.FileAttachment(zipDestinationPath, filepath.Base(zipDestinationPath))
}

func (inst *Controller) RestoreSnapshot(c *gin.Context) {
	if restoreStatus == model.Restoring {
		responseHandler(nil, errors.New("restoring snapshot"), c)
		return
	}
	restoreStatus = model.Restoring
	useGlobalUUID, err := utils.ToBool(c.Query("use_global_uuid"))
	if err != nil {
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	destinationFilePath := path.Join(config.Config.GetAbsTempDir(), file.Filename)
	err = c.SaveUploadedFile(file, destinationFilePath)
	if err != nil {
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	_, err = fileutils.Unzip(destinationFilePath, path.Join(config.Config.GetAbsTempDir(), ""), os.FileMode(inst.FileMode))
	if err != nil {
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	_ = os.RemoveAll(destinationFilePath)

	unzippedFolderPath := path.Join(config.Config.GetAbsTempDir(), utils.FileNameWithoutExtension(file.Filename))
	services, _ := fileutils.ListFiles(path.Join(unzippedFolderPath, systemFolder))
	inst.stopServices(services)
	err = utils.CopyDir(path.Join(unzippedFolderPath, systemFolder), systemPath)
	if err != nil {
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	err = utils.CopyDir(path.Join(unzippedFolderPath, dataFolder), config.Config.GetAbsDataDir())
	if err != nil {
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	_ = os.RemoveAll(unzippedFolderPath)
	if !useGlobalUUID {
		inst.updateDeviceInfo()
	}
	inst.enableAndRestartServices(services)
	message := model.Message{Message: "snapshot restored successfully"}
	restoreStatus = model.Restored
	responseHandler(message, err, c)
}

func (inst *Controller) SnapshotStatus(c *gin.Context) {
	responseHandler(model.SnapshotStatus{CreateStatus: createStatus, RestoreStatus: restoreStatus}, nil, c)
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
