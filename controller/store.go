package controller

import (
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/controller/response"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
	"net/http"
)

func getAppStoreBody(c *gin.Context) (dto *apps.Store, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetAppStores(c *gin.Context) {
	data, err := inst.DB.GetAppStores()
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) GetAppStore(c *gin.Context) {
	data, err := inst.DB.GetAppStore(c.Params.ByName("uuid"))
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) CreateAppStore(c *gin.Context) {
	var m *apps.Store
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.CreateAppStore(m)
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) UpdateAppStore(c *gin.Context) {
	body, _ := getAppStoreBody(c)
	data, err := inst.DB.UpdateAppStore(c.Params.ByName("uuid"), body)
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) DeleteAppStore(c *gin.Context) {
	data, err := inst.DB.DeleteAppStore(c.Params.ByName("uuid"))
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) DropAppStores(c *gin.Context) {
	data, err := inst.DB.DropAppStores()
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}
