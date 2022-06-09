package controller

import (
	"github.com/NubeIO/edge/service/apps/installer"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) UnInstallApp(c *gin.Context) {
	var m *installer.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.Installer.UnInstallApp(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) GetUnInstallProgress(c *gin.Context) {
	var m *installer.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.Installer.GetUnInstallProgress(m.AppName)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
