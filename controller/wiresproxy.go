package controller

import (
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/internaltoken"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func (inst *Controller) WiresProxy(c *gin.Context) { // eg http://0.0.0.0:1661/wires/api/nodes/values
	remote, err := Builder("0.0.0.0", 1665)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	internalToken := internaltoken.GetInternalToken(true)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
		req.Header.Set("Authorization", internalToken)
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
