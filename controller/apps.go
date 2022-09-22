package controller

import (
	"errors"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ListApps apps by listed in the installation (/data/rubix-service/apps/install)
func (inst *Controller) ListApps(c *gin.Context) {
	data, err := inst.EdgeApp.App.ListApps()
	responseHandler(data, err, c)
}

// ListAppsStatus get all the apps by listed in the installation (/data/rubix-service/apps/install) dir and then check the service
func (inst *Controller) ListAppsStatus(c *gin.Context) {
	data, err := inst.EdgeApp.App.ListAppsStatus()
	responseHandler(data, err, c)
}

// UploadApp uploads the build
func (inst *Controller) UploadApp(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		Name:                            c.Query("name"),
		Version:                         c.Query("version"),
		Product:                         c.Query("product"),
		Arch:                            c.Query("arch"),
		DoNotValidateArch:               c.Query("do_not_validate_arch") == "true",
		MoveExtractedFileToNameApp:      c.Query("move_extracted_file_to_name_app") == "true",
		MoveOneLevelInsideFileToOutside: c.Query("move_one_level_inside_file_to_outside") == "true",
		File:                            file,
	}
	data, err := inst.EdgeApp.App.UploadEdgeApp(m)
	responseHandler(data, err, c)
}

func (inst *Controller) UploadServiceFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		Name:    c.Query("name"),
		Version: c.Query("version"),
		File:    file,
	}
	data, err := inst.EdgeApp.App.UploadServiceFile(m)
	responseHandler(data, err, c)
}

func (inst *Controller) InstallService(c *gin.Context) {
	var m *installer.Install
	err := c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.EdgeApp.App.InstallService(m)
	responseHandler(data, err, c)
}

func (inst *Controller) UninstallApp(c *gin.Context) {
	deleteApp, _ := strconv.ParseBool(c.Query("delete"))
	name := c.Query("name")
	if name == "" {
		responseHandler(nil, errors.New("app_name can not be empty"), c)
		return
	}
	data, err := inst.EdgeApp.App.UninstallApp(name, deleteApp)
	responseHandler(data, err, c)
}
