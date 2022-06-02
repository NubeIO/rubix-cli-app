package controller

import (
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/controller/response"
	dbase "gthub.com/NubeIO/rubix-cli-app/database"
	"net/http"
)

func (inst *Controller) UnInstallApp(c *gin.Context) {
	var m *dbase.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.UnInstallApp(m)
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) GetUnInstallProgress(c *gin.Context) {
	var m *dbase.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.GetUnInstallProgress(m.AppName)
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}
