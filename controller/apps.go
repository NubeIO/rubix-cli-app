package controller

import (
	"errors"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-edge/models"
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
	appServiceMapping := models.AppServiceMapping{}
	err := c.ShouldBindJSON(&appServiceMapping)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.EdgeApp.App.ListAppsStatus(appServiceMapping.AppServiceMapping)
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
		Name:    c.Query("name"),
		Version: c.Query("version"),
		Product: c.Query("product"),
		Arch:    c.Query("arch"),
		File:    file,
	}
	data, err := inst.EdgeApp.App.UploadEdgeApp(m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
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
	serviceName := c.Query("service_name")
	if name != "" {
		responseHandler(nil, errors.New("name can not be empty"), c)
		return
	}
	if serviceName != "" {
		responseHandler(nil, errors.New("service_name can not be empty"), c)
		return
	}
	data, err := inst.EdgeApp.App.UninstallApp(name, serviceName, deleteApp)
	responseHandler(data, err, c)
}
