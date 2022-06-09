package controller

import (
	"github.com/NubeIO/edge/service/apps"
	"github.com/gin-gonic/gin"
)

func getAppsBody(c *gin.Context) (dto *apps.App, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetApps(c *gin.Context) {
	data, err := inst.DB.GetApps()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) AppStats(c *gin.Context) {
	body, err := getAppsBody(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.DB.AppStats(body)
	reposeHandler(data, err, c)
}

func (inst *Controller) GetApp(c *gin.Context) {
	data, err := inst.DB.GetApp(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) UpdateApp(c *gin.Context) {
	body, _ := getAppsBody(c)
	data, err := inst.DB.UpdateApp(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) DeleteApp(c *gin.Context) {
	data, err := inst.DB.DeleteApp(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)

}

func (inst *Controller) DropApps(c *gin.Context) {
	data, err := inst.DB.DropApps()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
