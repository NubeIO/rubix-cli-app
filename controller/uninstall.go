package controller

import (
	dbase "github.com/NubeIO/rubix-cli-app/database"
	"github.com/gin-gonic/gin"
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
