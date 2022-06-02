package client

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
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

func (response *Response) buildResponse(restyResp *resty.Response, err error) *Response {
	response.resty = restyResp
	var msg interface{}
	if err != nil {
		msg = err.Error()
	} else {
		asJson, err := response.AsJson()
		if err != nil {
			msg = restyResp.String()
		} else {
			msg = asJson
		}
	}
	response.StatusCode = restyResp.StatusCode()
	response.Message = msg
	return response
}

func (response *Response) AsString() string {
	return response.resty.String()
}

func (response *Response) GetError() interface{} {
	return response.resty.Error()
}

func (response *Response) GetStatus() int {
	return response.resty.StatusCode()
}

// AsJson return as body as blank interface
func (response *Response) AsJson() (interface{}, error) {
	var out interface{}
	err := json.Unmarshal(response.resty.Body(), &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (inst *Client) GetApps() (data []apps.App, response *Response) {
	path := fmt.Sprintf(Paths.Apps.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]apps.App{}).
		Get(path)
	return *resp.Result().(*[]apps.App), response.buildResponse(resp, err)
}
