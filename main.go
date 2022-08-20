package main

import (
	"fmt"
	"github.com/NubeIO/rubix-edge/service/system"
)

func main() {

	sys := system.New(&system.System{})
	fmt.Println(sys.GetHardwareTZ())
	//list, err := sys.GetTimeZoneList()
	//fmt.Println(err)
	//pprint.PrintJOSN(list)
	//cmd.Execute()

}
