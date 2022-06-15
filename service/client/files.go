package client

import (
	"fmt"
)

type UploadResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Destination string `json:"destination,omitempty"`
		File        string `json:"file,omitempty"`
		Size        string `json:"size,omitempty"`
		UploadTime  string `json:"upload_time,omitempty"`
	} `json:"data"`
}

func (inst *Client) UploadFile(file, destination string) (*UploadResponse, error) {
	path := fmt.Sprintf("/api/files/upload?destination=%s", destination)
	response := &UploadResponse{}
	resp, err := inst.Rest.R().
		SetResult(&UploadResponse{}).
		SetError(&UploadResponse{}).
		SetFile("file", file).
		Post(path)
	if err != nil {
		response.Code = 500
		response.Msg = err.Error()
		return response, err
	}
	response.Code = resp.StatusCode()
	if resp.IsSuccess() {
		return resp.Result().(*UploadResponse), nil
	}
	return resp.Error().(*UploadResponse), err
}
