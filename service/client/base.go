package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Rest *resty.Client
}

// New returns a new instance of the nube common apis
func New(url string, port int) *Client {
	rest := &Client{
		Rest: resty.New(),
	}
	if url == "" {
		url = "0.0.0.0"
	}
	if port == 0 {
		port = 1661
	}
	rest.Rest.SetBaseURL(fmt.Sprintf("http://%s:%d", url, port))
	return rest
}
