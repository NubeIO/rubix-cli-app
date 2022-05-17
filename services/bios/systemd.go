package bios

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
)

func build() {
	name := "aidans-service"
	user := "aidan"
	directory := "/home/aidan"
	execCmd := "/usr/bin/python3 something.py"

	write := builder.WriteFile{
		Write:    true,
		Path:     "/home/aidan",
		FileName: "test",
	}

	bld := &builder.SystemDBuilder{
		Name:      name,
		User:      user,
		Directory: directory,
		ExecCmd:   execCmd,
		WriteFile: write,
	}

	err := bld.Build()
	if err != nil {
		fmt.Println(err)
	}

}

func service() {
	timeOut := 5

	ctl.New(&ctl.Options{WorkDir: ""})
	opts := systemctl.Options{Timeout: timeOut}

	out, msg, err := systemctl.IsActive("mosquitto", opts)
	fmt.Println(out, msg)
	fmt.Println(err)

}
