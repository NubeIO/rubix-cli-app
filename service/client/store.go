package client

import (
	"github.com/go-resty/resty/v2"
)

type Path struct {
	Path string
}

var Paths = struct {
	Apps   Path
	Store  Path
	System Path
}{
	Apps:   Path{Path: "/api/apps"},
	Store:  Path{Path: "/api/store"},
	System: Path{Path: "/api/system"},
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
	resty      *resty.Response
}
