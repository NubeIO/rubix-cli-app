package controller

import (
	"github.com/NubeIO/lib-networking/networking"
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/controller/response"
	"gthub.com/NubeIO/rubix-cli-app/pkg/model"
	"gthub.com/NubeIO/rubix-cli-app/service/system"
	"net/http"
)

var nets = networking.New()

func (inst *Controller) GetIpSchema(c *gin.Context) {
	data := model.GetIpSchema()
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}

func (inst *Controller) Networking(c *gin.Context) {
	data, err := nets.GetNetworks()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) GetInterfacesNames(c *gin.Context) {
	data, err := nets.GetInterfacesNames()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) InternetIP(c *gin.Context) {
	data, err := nets.GetInternetIP()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) SetDHCP(c *gin.Context) {
	var m *system.IP
	err = c.ShouldBindJSON(&m)
	m.DHCP = true
	ip := system.NewIP(m)
	data, err := ip.SetDHCP()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) SetStaticIP(c *gin.Context) {
	var m *system.IP
	err = c.ShouldBindJSON(&m)
	ip := system.NewIP(m)
	data, err := ip.SetStaticIP()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
