package remote

import (
	"fmt"
	pprint "gthub.com/NubeIO/rubix-cli-app/pkg/helpers/print"
	"testing"
)

func TestAdmin_GetHardwareClock(t *testing.T) {
	host := &Admin{}
	run := New(host)
	run.ArchIsLinux()

	out, err := run.SystemTime()
	fmt.Println(err)
	pprint.PrintJOSN(out)
}
