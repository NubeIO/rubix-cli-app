package controller

import (
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetProduct(c *gin.Context) {
	data, err := inst.Rubix.App.GetProduct()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
