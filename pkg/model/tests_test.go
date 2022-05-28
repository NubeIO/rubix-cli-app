package model

import (
	"fmt"
	"gthub.com/NubeIO/rubix-cli-app/service/arch"
	"testing"
)

func TestProduct(t *testing.T) {

	err := arch.CheckProduct("RubixCompute")
	fmt.Println(err)
	if err != nil {
		//return
	}

}
