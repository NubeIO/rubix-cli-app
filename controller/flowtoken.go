package controller

import (
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) FlowToken(c *gin.Context) {
	sys := system.New(&system.System{})
	data, err := sys.GetFlowToken()
	if err != nil {
		responseHandler(data, err, c)
		return
	}
	responseHandler(data, err, c)
}
