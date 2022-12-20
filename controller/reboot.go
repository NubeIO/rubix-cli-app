package controller

import "github.com/gin-gonic/gin"

func (inst *Controller) RebootHost(c *gin.Context) {
	data, err := inst.System.RebootHost()
	responseHandler(data, err, c)
}
