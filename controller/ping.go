package controller

import "github.com/gin-gonic/gin"

func (inst *Controller) Ping(c *gin.Context) {
	message := Message{Message: "pong"}
	responseHandler(message, nil, c)
}
