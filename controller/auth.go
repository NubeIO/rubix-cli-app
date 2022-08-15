package controller

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (inst *Controller) HandleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorized := auth.Authorize(c.Request)
		if !authorized {
			c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized access"))
			return
		}
		c.Next()
		return
	}
}
