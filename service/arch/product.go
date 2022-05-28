package arch

import (
	"errors"
)

type ProductType int

//go:generate stringer -type=ProductType
const (
	RubixCompute ProductType = iota
	RubixComputeIO
	RubixCompute5
	Edge28
	Nuc
	None
)

func CheckProduct(s string) error {

	//cmd := &remote.Admin{}
	//run := remote.New(cmd)
	//_, host := run.DetectArch()
	//host.IsAMD64

	switch s {
	case RubixCompute.String():
		return nil
	case RubixComputeIO.String():
		return nil
	}
	return errors.New("invalid product type, try RubixCompute")

}
