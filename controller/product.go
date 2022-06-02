package controller

import (
	"github.com/NubeIO/lib-command/product"
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/controller/response"
	"net/http"
)

func (inst *Controller) GetProduct(c *gin.Context) {
	data, err := product.Get()
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}
