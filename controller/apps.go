package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddUploadApp
// upload the build
func (inst *Controller) AddUploadApp(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		Name:    c.Query("name"),
		Version: c.Query("version"),
		Product: c.Query("product"),
		Arch:    c.Query("arch"),
		File:    file,
	}
	data, err := inst.Rubix.App.AddUploadEdgeApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// UploadService
// upload the service file
func (inst *Controller) UploadService(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		Name:    c.Query("name"),
		Version: c.Query("version"),
		File:    file,
	}
	data, err := inst.Rubix.App.UploadServiceFile(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) InstallService(c *gin.Context) {
	var m *installer.Install
	err = c.ShouldBindJSON(&m)
	data, err := inst.Rubix.App.InstallService(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) UninstallService(c *gin.Context) {
	deleteApp, _ := strconv.ParseBool(c.Query("delete"))
	data, err := inst.Rubix.App.UninstallApp(c.Query("name"), c.Query("service"), deleteApp)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) RemoveApp(c *gin.Context) {
	err := inst.Rubix.App.RemoveApp(c.Query("name"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler("deleted app ok", nil, c)
}

func (inst *Controller) RemoveAppInstall(c *gin.Context) {
	err := inst.Rubix.App.RemoveAppInstall(c.Query("name"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler("deleted app ok", nil, c)
}
