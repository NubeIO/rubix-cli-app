package controller

import (
	"github.com/gin-gonic/gin"
	dbase "gthub.com/NubeIO/rubix-cli-app/database"
)

func (inst *Controller) UnInstallApp(c *gin.Context) {
	var m *dbase.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.UnInstallApp(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) GetUnInstallProgress(c *gin.Context) {
	var m *dbase.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.GetUnInstallProgress(m.AppName)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
