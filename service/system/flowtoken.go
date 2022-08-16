package system

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"strings"
)

type FlowToken struct {
	User  string `json:"user"`
	Token string `json:"token"`
}

// GetFlowToken do a system read to get the flow token
func (inst *System) GetFlowToken() (*FlowToken, error) {
	path := "/data/auth"
	fileName := "user.txt"
	files, err := fileutils.New().ListFiles(path)
	if err != nil {
		return nil, err
	}
	flowToken := &FlowToken{}
	if len(files) > 0 {
		for _, file := range files {
			if file == fileName {
				readFile, err := fileutils.New().ReadFile(fmt.Sprintf("%s/%s", path, fileName))
				if err != nil {
					return nil, err
				}
				parts := strings.Split(readFile, ":")
				if len(parts) == 2 {
					flowToken.User = parts[0]
					flowToken.Token = parts[1]
				}

			}

		}
	}
	return flowToken, err

}
