package product

import (
	"errors"
	"github.com/NubeIO/lib-dirs/dirs/jparse"
	log "github.com/sirupsen/logrus"
	"gthub.com/NubeIO/rubix-cli-app/pkg/config"
	"gthub.com/NubeIO/rubix-cli-app/service/remote"
)

var cmd = remote.New(&remote.Admin{})

type Product struct {
	Type     string `json:"type"`
	Version  string `json:"version"`
	Hardware string `json:"hardware"`
	Arch     string `json:"arch"` //amd64
}

func Get() (*Product, error) {
	return read()
}

func read() (*Product, error) {
	p := &Product{}
	j := jparse.New()
	var err error
	if readErr := j.ParseToData(config.ProductFilePath, p); readErr != nil {
		log.Errorln("read-product: read from json err", readErr.Error())
		err = readErr
	}
	_, _, model := cmd.DetectNubeProduct()
	if model == "" {
		archType := cmd.ArchCheck()
		if archType.Linux {
			model = "Linux"
		}
		if archType.Darwin {
			model = "Darwin"
		}
	}
	resp, err := cmd.DetectArch()
	if err != nil {
		return nil, err
	}
	p.Hardware = model
	p.Arch = resp.ArchModel
	return p, err
}

func CheckProduct(s string) (ProductType, error) {
	switch s {
	case RubixCompute.String():
		return RubixCompute, nil
	case RubixComputeIO.String():
		return RubixComputeIO, nil
	case Edge28.String():
		return Edge28, nil
	case AllLinux.String():
		return AllLinux, nil
	case RubixCompute5.String():
		return RubixCompute5, nil
	case Nuc.String():
		return Nuc, nil
	case Mac.String():
		return Mac, nil
	}

	return None, errors.New("invalid product type, try RubixCompute")

}
