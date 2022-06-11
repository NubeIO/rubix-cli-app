package controller

import (
	"github.com/NubeIO/edge/service/apps/installer"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) InstallApp(c *gin.Context) {

	var m *installer.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.Installer.InstallApp(m)
	if err != nil {
		reposeWithCode(404, data, err, c)
		return
	}
	reposeWithCode(202, data, err, c)
}

func (inst *Controller) GetInstallProgress(c *gin.Context) {
	var m *installer.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.Installer.GetInstallProgress(m.AppName)
	if err != nil {
		reposeWithCode(404, data, err, c)
		return
	}
	reposeWithCode(202, data, err, c)
}
