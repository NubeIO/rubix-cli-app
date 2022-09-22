package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
RESTORE A BACK-UP
*/

func (inst *Controller) RestoreBackup(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	takeBackup := c.Query("take_backup")
	_takeBackup, _ := strconv.ParseBool(takeBackup)
	m := &installer.RestoreBackup{
		TakeBackup: _takeBackup,
		File:       file,
	}
	data, err := inst.EdgeApp.RestoreBackup(m)
	responseHandler(data, err, c)
}

func (inst *Controller) RestoreAppBackup(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	takeBackup := c.Query("take_backup")
	_takeBackup, _ := strconv.ParseBool(takeBackup)
	m := &installer.RestoreBackup{
		AppName:    c.Query("app_name"),
		TakeBackup: _takeBackup,
		File:       file,
	}
	data, err := inst.EdgeApp.RestoreAppBackup(m)
	responseHandler(data, err, c)
}

/*
RUN A BACK-UP
*/

func (inst *Controller) FullBackUp(c *gin.Context) {
	data, err := inst.EdgeApp.FullBackUp(nil)
	responseHandler(data, err, c)
}

func (inst *Controller) BackupApp(c *gin.Context) {
	appName := c.Query("app_name")
	data, err := inst.EdgeApp.BackupApp(appName, nil)
	responseHandler(data, err, c)
}

/*
LIST BACK-UPS
*/

func (inst *Controller) ListFullBackups(c *gin.Context) {
	data, err := inst.EdgeApp.ListFullBackups()
	responseHandler(data, err, c)
}

func (inst *Controller) ListAppsBackups(c *gin.Context) {
	data, err := inst.EdgeApp.ListAppsBackups()
	responseHandler(data, err, c)
}

func (inst *Controller) ListAppBackups(c *gin.Context) {
	appName := c.Query("app_name")
	data, err := inst.EdgeApp.ListAppBackups(appName)
	responseHandler(data, err, c)
}

/*
DELETE BACK-UPS
*/

func (inst *Controller) WipeBackups(c *gin.Context) {
	data, err := inst.EdgeApp.WipeBackups()
	responseHandler(data, err, c)
}

func (inst *Controller) DeleteAllFullBackups(c *gin.Context) {
	data, err := inst.EdgeApp.DeleteAllFullBackups()
	responseHandler(data, err, c)
}

func (inst *Controller) DeleteAppsBackups(c *gin.Context) {
	data, err := inst.EdgeApp.DeleteAllAppsBackups()
	responseHandler(data, err, c)
}

func (inst *Controller) DeleteAllAppBackups(c *gin.Context) {
	appName := c.Query("name")
	data, err := inst.EdgeApp.DeleteAllAppBackups(appName)
	responseHandler(data, err, c)
}

func (inst *Controller) DeleteOneAppBackup(c *gin.Context) {
	appName := c.Query("name")
	zipFile := c.Query("zip_file")
	data, err := inst.EdgeApp.DeleteOneAppBackup(appName, zipFile)
	responseHandler(data, err, c)
}
