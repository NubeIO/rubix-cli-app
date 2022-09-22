package controller

import "github.com/gin-gonic/gin"

func (inst *Controller) GetNetworkInterfaces(c *gin.Context) {
	networks, err := nets.GetNetworks()
	responseHandler(networks, err, c)
}
