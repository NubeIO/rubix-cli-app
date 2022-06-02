package client

import (
	"fmt"
	"github.com/NubeIO/lib-date/datelib"
)

type Time struct {
	Code    int           `json:"code"`
	Message string        `json:"msg"`
	Data    *datelib.Time `json:"data"`
}

func (inst *Client) GetTime() (data *Time, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.System.Path, "time")
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&Time{}).
		Get(path)
	return resp.Result().(*Time), response.buildResponse(resp, err)
}
