package client

import (
	"fmt"
	"github.com/NubeIO/edge/service/apps"
	"github.com/NubeIO/edge/service/apps/installer"
)

type AppsResp struct {
	Code    int        `json:"code"`
	Message string     `json:"msg"`
	Data    []apps.App `json:"data"`
}

type AppResp struct {
	Code    int       `json:"code"`
	Message string    `json:"msg"`
	Data    *apps.App `json:"data"`
}

func (inst *Client) GetApps() (data []apps.App, response *Response) {
	path := fmt.Sprintf(Paths.Apps.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]apps.App{}).
		Get(path)
	return *resp.Result().(*[]apps.App), response.buildResponse(resp, err)
}

func (inst *Client) InstallApp(body *installer.App) (data *installer.InstallResponse, response *Response) {
	path := fmt.Sprintf(Paths.Apps.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&installer.InstallResponse{}).
		Post(path)
	return resp.Result().(*installer.InstallResponse), response.buildResponse(resp, err)
}

func (inst *Client) GetApp(uuid string) (data *AppResp, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Apps.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&AppResp{}).
		Get(path)
	return resp.Result().(*AppResp), response.buildResponse(resp, err)
}

func (inst *Client) UpdateApp(uuid string, body *apps.App) (data *apps.App, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Apps.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&apps.App{}).
		Patch(path)
	return resp.Result().(*apps.App), response.buildResponse(resp, err)
}

func (inst *Client) DeleteApp(uuid string) (response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Apps.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		Delete(path)
	return response.buildResponse(resp, err)
}
