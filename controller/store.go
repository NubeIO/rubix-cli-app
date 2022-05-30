package controller

import (
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

func getAppImageBody(ctx *gin.Context) (dto *apps.Store, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetAppImages(c *gin.Context) {
	data, err := inst.DB.GetAppImages()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) GetAppImage(c *gin.Context) {
	data, err := inst.DB.GetAppImage(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) CreateAppImage(c *gin.Context) {
	var m *apps.Store
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.CreateAppImage(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) UpdateAppImage(c *gin.Context) {
	body, _ := getAppImageBody(c)
	data, err := inst.DB.UpdateAppImage(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) DeleteAppImage(c *gin.Context) {
	q, err := inst.DB.DeleteAppImage(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (inst *Controller) DropAppImages(c *gin.Context) {
	data, err := inst.DB.DropAppImages()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}
