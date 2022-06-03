package controller

import (
	"github.com/NubeIO/rubix-cli-app/service/apps"
	"github.com/gin-gonic/gin"
)

func getAppStoreBody(c *gin.Context) (dto *apps.Store, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetAppStores(c *gin.Context) {
	data, err := inst.DB.GetAppStores()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) GetAppStore(c *gin.Context) {
	data, err := inst.DB.GetAppStore(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) CreateAppStore(c *gin.Context) {
	var m *apps.Store
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.CreateAppStore(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) UpdateAppStore(c *gin.Context) {
	body, _ := getAppStoreBody(c)
	data, err := inst.DB.UpdateAppStore(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) DeleteAppStore(c *gin.Context) {
	data, err := inst.DB.DeleteAppStore(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) DropAppStores(c *gin.Context) {
	data, err := inst.DB.DropAppStores()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
