package controller

import (
	"github.com/NubeIO/lib-date/datectl"
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) SystemTime(c *gin.Context) {
	data := inst.System.SystemTime()
	reposeHandler(data, nil, c)
}

func (inst *Controller) GenerateTimeSyncConfig(c *gin.Context) {
	var m *datectl.TimeSyncConfig
	err := c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data := inst.System.GenerateTimeSyncConfig(m)
	reposeHandler(data, nil, c)
}

func (inst *Controller) GetHardwareTZ(c *gin.Context) {
	data, err := inst.System.GetHardwareTZ()
	reposeHandler(data, err, c)
}

func (inst *Controller) GetTimeZoneList(c *gin.Context) {
	data, err := inst.System.GetTimeZoneList()
	reposeHandler(data, err, c)
}

func (inst *Controller) UpdateTimezone(c *gin.Context) {
	var m system.DateBody
	err := c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.System.UpdateTimezone(m)
	reposeHandler(data, err, c)
}

func (inst *Controller) SetSystemTime(c *gin.Context) {
	var m system.DateBody
	err := c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.System.SetSystemTime(m)
	reposeHandler(data, err, c)
}
