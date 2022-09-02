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
	data, err := inst.EdgeApp.FullBackUp()
	responseHandler(data, err, c)
}

func (inst *Controller) BackupApp(c *gin.Context) {
	appName := c.Query("app_name")
	data, err := inst.EdgeApp.BackupApp(appName)
	responseHandler(data, err, c)
}

/*
LIST BACK-UPS
*/

func (inst *Controller) ListFullBackups(c *gin.Context) {
	data, err := inst.EdgeApp.ListFullBackups()
	responseHandler(data, err, c)
}

func (inst *Controller) ListAppBackupsDirs(c *gin.Context) {
	data, err := inst.EdgeApp.ListAppBackupsDirs()
	responseHandler(data, err, c)
}

func (inst *Controller) ListBackupsByApp(c *gin.Context) {
	appName := c.Query("app_name")
	data, err := inst.EdgeApp.ListBackupsByApp(appName)
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

func (inst *Controller) DeleteAllAppBackups(c *gin.Context) {
	data, err := inst.EdgeApp.DeleteAllAppBackups()
	responseHandler(data, err, c)
}

func (inst *Controller) DeleteAppAllBackUpByName(c *gin.Context) {
	appName := c.Query("name")
	data, err := inst.EdgeApp.DeleteAppAllBackUpByName(appName)
	responseHandler(data, err, c)
}

func (inst *Controller) DeleteAppOneBackUpByName(c *gin.Context) {
	appName := c.Query("name")
	folder := c.Query("folder")
	data, err := inst.EdgeApp.DeleteAppOneBackUpByName(appName, folder)
	responseHandler(data, err, c)
}
