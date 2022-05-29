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
		if cmd.ArchIsLinux() {
			model = "linux"
		}
	}
	p.Hardware = model
	return p, err
}

func AutoCheckProduct(s string) error {
	switch s {
	case RubixCompute.String():
		return nil
	case RubixComputeIO.String():
		return nil
	}
	return errors.New("invalid product type, try RubixCompute")

}
