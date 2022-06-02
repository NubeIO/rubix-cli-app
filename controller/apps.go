package controller

import (
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/controller/response"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
	"net/http"
)

func getAppsBody(c *gin.Context) (dto *apps.App, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetApps(c *gin.Context) {
	data, err := inst.DB.GetApps()
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) AppStats(c *gin.Context) {
	body, err := getAppsBody(c)
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	data, err := inst.DB.AppStats(body)
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) GetApp(c *gin.Context) {
	data, err := inst.DB.GetApp(c.Params.ByName("uuid"))
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) UpdateApp(c *gin.Context) {
	body, _ := getAppsBody(c)
	data, err := inst.DB.UpdateApp(c.Params.ByName("uuid"), body)
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) DeleteApp(c *gin.Context) {
	data, err := inst.DB.DeleteApp(c.Params.ByName("uuid"))
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)

}

func (inst *Controller) DropApps(c *gin.Context) {
	data, err := inst.DB.DropApps()
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}
