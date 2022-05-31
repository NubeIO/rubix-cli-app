package controller

import (
	"github.com/NubeIO/lib-command/product"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetProduct(c *gin.Context) {
	data, err := product.Get()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}
