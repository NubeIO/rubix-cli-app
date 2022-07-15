package rubix

import (
	"fmt"
	"testing"
)

func Test_checkVersion(t *testing.T) {

	checkVersion("v0.2.3")

	err := makeAppInstallDir("/home/aidan/test6")
	fmt.Println(err)
	if err != nil {
		return
	}
}
