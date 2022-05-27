package model

import (
	"fmt"
	"testing"
)

func TestProduct(t *testing.T) {

	err := CheckProduct("RubixCompute")
	fmt.Println(err)
	if err != nil {
		//return
	}

}
