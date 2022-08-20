package system

import (
	"errors"
	"os/exec"
	"strings"
)

func (inst *System) GetHardwareTZ() (string, error) {
	cmd := exec.Command("cat", "/etc/timezone")
	output, err := cmd.Output()
	cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return "", err
	}
	out := strings.Split(string(output), "\n")
	if len(out) >= 0 {
		return out[0], nil
	} else {
		return "", errors.New("failed to find timezone")
	}
}

func (inst *System) GetTimeZoneList() ([]string, error) {
	cmd := exec.Command("timedatectl", "list-timezones")
	output, err := cmd.Output()
	cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return nil, err
	}
	var out []string
	list := strings.Split(string(output), "\n")
	for _, s := range list {
		if s != "" {
			out = append(out, s)
		}
	}
	return out, nil
}

// UpdateTimezone sets the current machine's timezone to the given timezone
func (inst *System) UpdateTimezone(newZone string) error {
	list, err := inst.GetTimeZoneList()
	if err != nil {
		return err
	}
	var matchZone bool
	for _, zone := range list {
		if zone == newZone {
			matchZone = true
		}
	}
	if !matchZone {
		return errors.New("incorrect zone passed in try, Australia/Sydney")
	}
	cmd := exec.Command("timedatectl", "set-timezones", strings.TrimSpace(newZone))
	output, err := cmd.Output()
	cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return err
	}

	return nil
}

//
//func (inst *System) SetSystemTime(date string) error {
//	layout := "Mon, 02 Jan 2006 15:04:05 MST"
//	// parse time
//	t, err := time.Parse(layout, date)
//
//	if err != nil {
//		return fmt.Errorf("could not parse date  %s", err)
//	}
//	// set system time
//	inst.CMD.Commands = Builder("date", "-s", fmt.Sprintf("@%d", t.Unix()))
//	res := inst.CMD.RunCommand()
//	if res.Err != nil {
//		return res.Err
//	}
//	return nil
//}
//
//// NTPDisable timedatectl set-ntp false
//func (inst *System) NTPDisable() error {
//	// set system time
//	inst.CMD.Commands = Builder("timedatectl", "-set-ntp", "false")
//	res := inst.CMD.RunCommand()
//	if res.Err != nil {
//		return res.Err
//	}
//	return nil
//}
//
//// NTPEnable timedatectl set-ntp true
//func (inst *System) NTPEnable() error {
//	// set system time
//	inst.CMD.Commands = Builder("timedatectl", "-set-ntp", "true")
//	res := inst.CMD.RunCommand()
//	if res.Err != nil {
//		return res.Err
//	}
//	return nil
//}
