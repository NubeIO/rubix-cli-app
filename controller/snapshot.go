package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-edge/model"
	"github.com/NubeIO/rubix-edge/pkg/config"
	"github.com/NubeIO/rubix-edge/service/clients/bioscli"
	"github.com/NubeIO/rubix-edge/utils"
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
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
	log.Info("creating snapshot...")
	if createStatus == model.Creating {
		err := errors.New("snapshot creation process is in progress")
		log.Error(err)
		responseHandler(nil, err, c)
		return
	}
	createStatus = model.Creating
	deviceInfo, err := inst.RubixRegistry.GetDeviceInfo()
	if err != nil {
		log.Error(err)
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	clientName := strings.Replace(deviceInfo.ClientName, "/", "", -1)
	siteName := strings.Replace(deviceInfo.SiteName, "/", "", -1)
	deviceName := strings.Replace(deviceInfo.DeviceName, "/", "", -1)
	if clientName == "" || clientName == "-" {
		clientName = "na"
	}
	if siteName == "" || siteName == "-" {
		siteName = "na"
	}
	if deviceName == "" || deviceName == "-" {
		deviceName = "na"
	}
	filePrefix := fmt.Sprintf("%s-%s-%s", clientName, siteName, deviceName)
	previousFiles, _ := filepath.Glob(path.Join(config.Config.GetAbsTempDir(), fmt.Sprintf("%s*", filePrefix)))
	utils.DeleteFiles(previousFiles, config.Config.GetAbsTempDir())

	biosClient := bioscli.NewLocalBiosClient()
	arch, err := biosClient.GetArch()
	if err != nil {
		log.Error(err)
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}

	destinationPath := fmt.Sprintf("%s/%s_%s_%s", config.Config.GetAbsTempDir(), filePrefix,
		time.Now().UTC().Format("20060102T150405"), arch.Arch)
	absDataFolder := path.Join(destinationPath, dataFolder)
	err = os.MkdirAll(absDataFolder, os.FileMode(inst.FileMode)) // create empty folder even we don't have content
	if err != nil {
		log.Error(err)
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	_ = utils.CopyDir(config.Config.GetSnapshotDir(), absDataFolder, "", 0)

	systemFiles, err := filepath.Glob(path.Join(systemPath, "nubeio-*"))
	if err != nil {
		log.Error(err)
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	absSystemFolder := path.Join(destinationPath, systemFolder)
	err = os.MkdirAll(absSystemFolder, os.FileMode(inst.FileMode)) // create empty folder even we don't have content
	if err != nil {
		log.Error(err)
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	utils.CopyFiles(systemFiles, absSystemFolder)

	zipDestinationPath := destinationPath + ".zip"
	log.Infof("zipping snapshot: %s...", zipDestinationPath)
	err = fileutils.RecursiveZip(destinationPath, zipDestinationPath)
	if err != nil {
		log.Error(err)
		createStatus = model.CreateFailed
		responseHandler(nil, err, c)
		return
	}
	_ = os.RemoveAll(destinationPath)
	createStatus = model.Created
	log.Info("sending snapshot data...")
	c.FileAttachment(zipDestinationPath, filepath.Base(zipDestinationPath))
}

func (inst *Controller) RestoreSnapshot(c *gin.Context) {
	log.Info("restoring snapshot...")
	if restoreStatus == model.Restoring {
		err := errors.New("snapshot restoring process is in progress")
		log.Error(err)
		responseHandler(nil, err, c)
		return
	}
	log.Info("receiving file data...")
	restoreStatus = model.Restoring
	file, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}

	biosClient := bioscli.NewLocalBiosClient()
	arch, err := biosClient.GetArch()
	if err != nil {
		log.Error(err)
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}

	fileParts := strings.Split(file.Filename, "_")
	archParts := fileParts[len(fileParts)-1]
	archFromSnapshot := strings.Split(archParts, ".")[0]
	if archFromSnapshot != arch.Arch {
		restoreStatus = model.RestoreFailed
		err = errors.New(
			fmt.Sprintf("arch mismatch: snapshot arch is %s & device arch is %s", archFromSnapshot, arch.Arch))
		log.Error(err)
		responseHandler(nil, err, c)
		return
	}

	log.Info("saving received file data...")
	destinationFilePath := path.Join(config.Config.GetAbsTempDir(), file.Filename)
	err = c.SaveUploadedFile(file, destinationFilePath)
	if err != nil {
		log.Error(err)
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	log.Info("unzipping file...")
	_, err = fileutils.Unzip(destinationFilePath, config.Config.GetAbsTempDir(), os.FileMode(inst.FileMode))
	if err != nil {
		log.Error(err)
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
		err = utils.CopyDir(path.Join(unzippedFolderPath, systemFolder), systemPath, "", 0)
		if err != nil {
			log.Error(err)
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
			log.Error(err)
			restoreStatus = model.RestoreFailed
			responseHandler(nil, err, c)
			return
		}
		err = inst.retainGlobalUUID(deviceInfo.GlobalUUID, rubixRegistryFile)
		if err != nil {
			log.Error(err)
			restoreStatus = model.RestoreFailed
			responseHandler(nil, err, c)
			return
		}
	}

	err = utils.DeleteDir(path.Join(unzippedFolderPath, dataFolder), "", 0)
	if err != nil {
		log.Error(err)
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	err = utils.CopyDir(path.Join(unzippedFolderPath, dataFolder), config.Config.GetSnapshotDir(), "", 0)
	if err != nil {
		restoreStatus = model.RestoreFailed
		responseHandler(nil, err, c)
		return
	}
	err = os.RemoveAll(unzippedFolderPath)
	if err != nil {
		log.Errorf("failed to remove file %s", unzippedFolderPath)
	}
	if copySystemFiles {
		err = inst.SystemCtl.DaemonReload()
		if err != nil {
			log.Error(err)
			restoreStatus = model.RestoreFailed
			responseHandler(nil, err, c)
			return
		}
		inst.enableAndRestartServices(services)
	}
	log.Info("snapshot is restored")
	message := model.Message{Message: "snapshot is restored successfully"}
	restoreStatus = model.Restored
	responseHandler(message, err, c)
}

func (inst *Controller) SnapshotStatus(c *gin.Context) {
	responseHandler(model.SnapshotStatus{CreateStatus: createStatus, RestoreStatus: restoreStatus}, nil, c)
}

func (inst *Controller) stopServices(services []string) {
	log.Info("stopping services...")
	var wg sync.WaitGroup
	for _, service := range services {
		wg.Add(1)
		go func(service string) {
			defer wg.Done()
			if !utils.Contains(utils.ExcludedServices, service) {
				err := inst.SystemCtl.Stop(service)
				if err != nil {
					log.Errorf("failed to stop service %s", service)
				}
			}
		}(service)
	}
	wg.Wait()
}

func (inst *Controller) enableAndRestartServices(services []string) {
	log.Info("enabling & restarting services")
	var wg sync.WaitGroup
	for _, service := range services {
		wg.Add(1)
		go func(service string) {
			defer wg.Done()
			if !utils.Contains(utils.ExcludedServices, service) {
				err := inst.SystemCtl.Enable(service)
				if err != nil {
					log.Errorf("failed to enable service %s", service)
				}
				err = inst.SystemCtl.Restart(service)
				if err != nil {
					log.Errorf("failed to restart service %s", service)
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
