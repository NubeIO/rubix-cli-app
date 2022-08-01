package controller

import (
	"github.com/NubeIO/rubix-edge/service/apps"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) CtlAction(c *gin.Context) {
	var m *apps.CtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.CtlAction(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) CtlStatus(c *gin.Context) {
	var m *apps.CtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.CtlStatus(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) ServiceMassAction(c *gin.Context) {
	var m *apps.CtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.ServiceMassAction(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) ServiceMassCheck(c *gin.Context) {
	var m *apps.CtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.ServiceMassCheck(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
