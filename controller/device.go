package controller

import "github.com/gin-gonic/gin"

func (inst *Controller) GetDeviceInfo(c *gin.Context) {
	data, err := inst.DB.GetDeviceInfo(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) AddDeviceInfo(c *gin.Context) {
	body, _ := getDeviceBody(c)
	data, err := inst.DB.AddDeviceInfo(body)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) UpdateDeviceInfo(c *gin.Context) {
	body, _ := getDeviceBody(c)
	data, err := inst.DB.UpdateDeviceInfo(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
