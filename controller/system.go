package controller

import (
	"github.com/NubeIO/lib-date/datelib"
	"github.com/NubeIO/lib-networking/networking"
	"github.com/gin-gonic/gin"
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

	data, err := datelib.New(&datelib.Date{}).SystemTime()
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
