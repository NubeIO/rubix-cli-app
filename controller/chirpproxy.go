package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func (inst *Controller) ChirpProxy(c *gin.Context) { // eg http://0.0.0.0:8080/chrip/api/organizations?limit=10
	remote, err := Builder("0.0.0.0", 8080)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
		req.Header.Set("Grpc-Metadata-Authorization", c.GetHeader("cs-token")) // pass in a header with the chirp-stack auth token
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
