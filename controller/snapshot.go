package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-edge/model"
	"github.com/NubeIO/rubix-edge/pkg/config"
	"github.com/NubeIO/rubix-edge/utils"
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
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
		responseHandler(nil, errors.New("snapshot creation process is in progress"), c)
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
	absDataFolder := path.Join(destinationPath, dataFolder)
	err = os.MkdirAll(absDataFolder, os.FileMode(inst.FileMode)) // create empty folder even we don't have content
	if err != nil {
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	_ = utils.CopyDir(config.Config.GetSnapshotDir(), absDataFolder, 0)

	systemFiles, err := filepath.Glob(path.Join(systemPath, "nubeio-*"))
	if err != nil {
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	absSystemFolder := path.Join(destinationPath, systemFolder)
	err = os.MkdirAll(absSystemFolder, os.FileMode(inst.FileMode)) // create empty folder even we don't have content
	if err != nil {
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	utils.CopyFiles(systemFiles, absSystemFolder)

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
		responseHandler(nil, errors.New("snapshot restoring process is in progress"), c)
		return
	}
	restoreStatus = model.Restoring
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
	_, err = fileutils.Unzip(destinationFilePath, config.Config.GetAbsTempDir(), os.FileMode(inst.FileMode))
	if err != nil {
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	_ = os.RemoveAll(destinationFilePath)

	unzippedFolderPath := path.Join(config.Config.GetAbsTempDir(), utils.FileNameWithoutExtension(file.Filename))

	copySystemFiles := true // for example in macOS, we don't have systemd file & so to prevent that failure
	services := make([]string, 0)
	if _, err := os.Stat(systemPath); errors.Is(err, os.ErrNotExist) {
		copySystemFiles = false
	}
	if copySystemFiles {
		services, _ = fileutils.ListFiles(path.Join(unzippedFolderPath, systemFolder))
		inst.stopServices(services)
		err = utils.CopyDir(path.Join(unzippedFolderPath, systemFolder), systemPath, 0)
		if err != nil {
			restoreStatus = model.RestoreFailed
			responseHandler(nil, err, c)
			return
		}
	}
	rubixRegistryFile := path.Join(unzippedFolderPath, inst.RubixRegistry.RubixRegistryDeviceInfoFile)
	rubixRegistryFileExist := false
	if _, err = os.Stat(rubixRegistryFile); !errors.Is(err, os.ErrNotExist) {
		rubixRegistryFileExist = true
	}
	if rubixRegistryFileExist {
		deviceInfo, err := inst.RubixRegistry.GetDeviceInfo()
		if err != nil {
			restoreStatus = model.RestoreFailed
			responseHandler(nil, err, c)
			return
		}
		err = inst.retainGlobalUUID(deviceInfo.GlobalUUID, rubixRegistryFile)
		if err != nil {
			restoreStatus = model.RestoreFailed
			responseHandler(nil, err, c)
			return
		}
	}

	err = utils.CopyDir(path.Join(unzippedFolderPath, dataFolder), config.Config.GetSnapshotDir(), 0)
	if err != nil {
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	_ = os.RemoveAll(unzippedFolderPath)
	if copySystemFiles {
		err = inst.SystemCtl.DaemonReload()
		if err != nil {
			restoreStatus = model.RestoreFailed
			responseHandler(nil, err, c)
			return
		}
		inst.enableAndRestartServices(services)
	}
	message := model.Message{Message: "snapshot is restored successfully"}
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
		go func(service string) {
			defer wg.Done()
			if service != "nubeio-rubix-edge.service" && service != "nubeio-rubix-assist.service" {
				err := inst.SystemCtl.Stop(service)
				if err != nil {
					log.Errorf("err: %s", err.Error())
				}
			}
		}(service)
	}
	wg.Wait()
}

func (inst *Controller) enableAndRestartServices(services []string) {
	var wg sync.WaitGroup
	for _, service := range services {
		wg.Add(1)
		go func(service string) {
			defer wg.Done()
			if service != "nubeio-rubix-edge.service" && service != "nubeio-rubix-assist.service" {
				err := inst.SystemCtl.Enable(service)
				if err != nil {
					log.Errorf("err: %s", err.Error())
				}
				err = inst.SystemCtl.Restart(service)
				if err != nil {
					log.Errorf("err: %s", err.Error())
				}
			}
		}(service)
	}
	wg.Wait()
}

func (inst *Controller) retainGlobalUUID(globalUUID, rubixRegistryFile string) error {
	content, err := fileutils.ReadFile(rubixRegistryFile)
	if err != nil {
		return err
	}
	deviceInfoDefault := rubixregistry.DeviceInfoDefault{}
	err = json.Unmarshal([]byte(content), &deviceInfoDefault)
	if err != nil {
		return err
	}
	deviceInfoDefault.DeviceInfoFirstRecord.DeviceInfo.GlobalUUID = globalUUID
	deviceInfoDefaultRaw, err := json.Marshal(deviceInfoDefault)
	if err != nil {
		return err
	}
	err = os.WriteFile(rubixRegistryFile, deviceInfoDefaultRaw, os.FileMode(inst.FileMode))
	if err != nil {
		return err
	}
	return nil
}
