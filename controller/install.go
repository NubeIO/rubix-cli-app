package controller

import (
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/controller/response"
	dbase "gthub.com/NubeIO/rubix-cli-app/database"
	"net/http"
)

func (inst *Controller) InstallApp(c *gin.Context) {
	var m *dbase.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.InstallApp(m)
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) GetInstallProgress(c *gin.Context) {
	var m *dbase.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.GetInstallProgress(m.AppName)
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}
