package model

import (
	"fmt"
	"gthub.com/NubeIO/rubix-cli-app/service/product"
	"testing"
)

func TestProduct(t *testing.T) {

	err := product.CheckProduct("RubixCompute")
	fmt.Println(err)
	if err != nil {
		//return
	}

}
