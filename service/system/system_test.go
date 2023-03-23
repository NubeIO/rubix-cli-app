package system

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-edge/pkg/helpers/print"

	"testing"
)

func TestSystem_DiscUsage(t *testing.T) {
	sys := New(&System{})
	r, err := sys.GetSystem()
	if err != nil {
		fmt.Println(err)
	}
	pprint.PrintJOSN(r)
}
