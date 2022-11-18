package controller

import (
	"github.com/NubeIO/rubix-edge/model"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) Ping(c *gin.Context) {
	message := model.Message{Message: "pong"}
	responseHandler(message, nil, c)
}
