package bioscli

import (
	"fmt"
	"github.com/NubeIO/rubix-edge/model"
	"github.com/NubeIO/rubix-edge/nresty"
)

func (inst *BiosClient) GetArch() (*model.Arch, error) {
	url := fmt.Sprintf("/api/system/arch")
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Arch{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Arch), nil
}
