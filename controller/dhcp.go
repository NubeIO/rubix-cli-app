package controller

import (
	"github.com/NubeIO/lib-dhcpd/dhcpd"
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) DHCPPortExists(c *gin.Context) {
	var m system.NetworkingBody
	err := c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.System.DHCPPortExists(m)
	responseHandler(data, err, c)
}

func (inst *Controller) DHCPSetAsAuto(c *gin.Context) {
	var m system.NetworkingBody
	err := c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.System.DHCPSetAsAuto(m)
	responseHandler(data, err, c)
}

func (inst *Controller) DHCPSetStaticIP(c *gin.Context) {
	var m *dhcpd.SetStaticIP
	err := c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.System.DHCPSetStaticIP(m)
	responseHandler(data, err, c)
}
