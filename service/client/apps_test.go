package client

import (
	"fmt"
	"testing"
)

func TestHost(*testing.T) {

	client := New("0.0.0.0", 8090)

	data, res := client.GetTime()

	fmt.Println(data.Data.DateFormatLocal, res.GetStatus())

}
