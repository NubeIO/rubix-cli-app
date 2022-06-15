package controller

import (
	"github.com/NubeIO/edge/service/system"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) RunScanner(c *gin.Context) {
	var m *system.Scanner
	err = c.ShouldBindJSON(&m)
	data, err := system.RunScanner(m)
	if err != nil {
		reposeWithCode(404, data, err, c)
		return
	}
	reposeWithCode(202, data, err, c)
}
