package controller

import (
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/pkg/model"
)

func getAppsBody(ctx *gin.Context) (dto *model.Apps, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetApps(c *gin.Context) {
	hosts, err := inst.DB.GetUsers()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(hosts, err, c)
}

func (inst *Controller) GetApp(c *gin.Context) {
	host, err := inst.DB.GetApp(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (inst *Controller) CreateApp(c *gin.Context) {
	var m *model.Apps
	err = c.ShouldBindJSON(&m)
	host, err := inst.DB.CreateApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (inst *Controller) UpdateApp(c *gin.Context) {
	body, _ := getAppsBody(c)
	host, err := inst.DB.UpdateApp(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
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
	host, err := inst.DB.DropApps()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}
