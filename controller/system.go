package controller

import (
	"github.com/NubeIO/lib-networking/networking"
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/service/remote"
)

func (inst *Controller) Networking(c *gin.Context) {
	nets := networking.NewNets()
	data, err := nets.GetNetworks()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) GetInterfacesNames(c *gin.Context) {
	nets := networking.NewNets()
	data, err := nets.GetInterfacesNames()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) HostTime(c *gin.Context) {
	host := &remote.Admin{}
	run := remote.New(host)
	data, err := run.SystemTime()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) InternetIP(c *gin.Context) {
	nets := networking.NewNets()
	data, err := nets.GetInternetIP()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}
