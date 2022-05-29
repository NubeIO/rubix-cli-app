package controller

import (
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

func getAppsBody(ctx *gin.Context) (dto *apps.Store, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetApps(c *gin.Context) {
	data, err := inst.DB.GetApps()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) GetApp(c *gin.Context) {
	data, err := inst.DB.GetApp(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) CreateApp(c *gin.Context) {
	var m *apps.Store
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.CreateApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) UpdateApp(c *gin.Context) {
	body, _ := getAppsBody(c)
	data, err := inst.DB.UpdateApp(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) DeleteApp(c *gin.Context) {
	q, err := inst.DB.DeleteApp(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (inst *Controller) DropApps(c *gin.Context) {
	data, err := inst.DB.DropApps()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}
