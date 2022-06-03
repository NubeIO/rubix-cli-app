package client

import (
	"fmt"
	dbase "github.com/NubeIO/rubix-cli-app/database"
	"github.com/NubeIO/rubix-cli-app/service/apps"
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

func (inst *Client) InstallApp(body *dbase.App) (data *dbase.InstallResponse, response *Response) {
	path := fmt.Sprintf(Paths.Apps.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&dbase.InstallResponse{}).
		Post(path)
	return resp.Result().(*dbase.InstallResponse), response.buildResponse(resp, err)
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
