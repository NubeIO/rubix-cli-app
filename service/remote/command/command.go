package command

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type Command struct {
	ShellToUse string
	SetPath    string
	Commands   []string
}

type Response struct {
	Ok      bool
	Out     string
	OutByte []byte
	Err     error
}

func Builder(args ...string) []string {
	return args
}

//TODO need to make its easier to build commands as the remote.SSH only takes in one arg as a string

func (inst *Command) RunCommand() (res *Response) {
	res = &Response{}
	if len(inst.Commands) <= 0 {
		res.Err = fmt.Errorf("no command provided")
		return res
	}
	shell := inst.ShellToUse //bash -c, "/usr/bin/ls"
	if shell == "" {
		shell = inst.Commands[0]
	}

	log.Infoln("CMD to run", exec.Command(shell, inst.Commands[1:]...).String())
	cmd := exec.Command(shell, inst.Commands[1:]...)
	cmd.Dir = inst.SetPath
	output, err := cmd.Output()
	outAsString := strings.TrimRight(string(output), "\n")
	if err != nil {
		log.Infoln("cmd", err)
		res.Err = err
		return res
	}
	res.Out = outAsString
	res.Ok = true
	res.OutByte = output
	return res
}
