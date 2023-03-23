package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Rican7/conjson"
	"github.com/Rican7/conjson/transform"
	"github.com/dhamith93/systats"
	"strings"
)

func (inst *System) GetSystem() (systats.System, error) {
	return inst.syStats.GetSystem()
}

func (inst *System) DiscUsage() ([]systats.Disk, error) {
	return inst.syStats.GetDisks()
}

type DiskUsage struct {
	Size      string `json:"size"`
	Used      string `json:"used"`
	Available string `json:"available"`
	Usage     string `json:"usage"`
}

type Disk struct {
	FileSystem string    `json:"file_system"`
	Type       string    `json:"type"`
	MountedOn  string    `json:"mounted_on"`
	Usage      DiskUsage `json:"usage"`
}

func (inst *System) DiscUsagePretty() ([]Disk, error) {
	var out []Disk
	disks, err := inst.syStats.GetDisks()
	if err != nil {
		return nil, err
	}
	for _, disk := range disks {
		newDisk := Disk{
			FileSystem: disk.FileSystem,
			Type:       disk.Type,
			MountedOn:  disk.MountedOn,
			Usage: DiskUsage{
				Size:      bytePretty(disk.Usage.Size),
				Used:      bytePretty(disk.Usage.Used),
				Available: bytePretty(disk.Usage.Available),
				Usage:     disk.Usage.Usage,
			},
		}
		out = append(out, newDisk)
	}
	return out, nil
}

type Memory struct {
	Stats     systats.Memory `json:"stats"`
	Available string
	Free      string
	Used      string
	Total     string
	Unit      string
}

type TopProcesses struct {
	Count int    `json:"count"`
	Sort  string `json:"sort"`
}

func (inst *System) GetTopProcesses(body TopProcesses) ([]systats.Process, error) {
	count := body.Count
	sort := body.Sort
	if count == 0 {
		count = 1
	}
	var correctSort bool
	if sort == "" {
		sort = "memory"
	}
	if sort != "memory" {
		correctSort = true
	}
	if sort != "cpu" {
		correctSort = true
	}
	if !correctSort {
		return nil, errors.New("incorrect sort type try: cpu, memory")
	}
	return inst.syStats.GetTopProcesses(count, sort)
}

func (inst *System) GetMemory() (systats.Memory, error) {
	return inst.syStats.GetMemory(systats.Megabyte)
}

func (inst *System) GetSwap() (systats.Swap, error) {
	return inst.syStats.GetSwap(systats.Megabyte)
}

func ConvertJSONKeys(s interface{}) json.Marshaler {
	return conjson.NewMarshaler(s, transform.ConventionalKeys())
}

const (
	BYTE     = 1.0
	KILOBYTE = 1024 * BYTE
	MEGABYTE = 1024 * KILOBYTE
	GIGABYTE = 1024 * MEGABYTE
	TERABYTE = 1024 * GIGABYTE
)

func bytePretty(size uint64) string {
	unit := ""
	value := float32(size)

	switch {
	case size >= TERABYTE:
		unit = "T"
		value = value / TERABYTE
	case size >= GIGABYTE:
		unit = "G"
		value = value / GIGABYTE
	case size >= MEGABYTE:
		unit = "M"
		value = value / MEGABYTE
	case size >= KILOBYTE:
		unit = "K"
		value = value / KILOBYTE
	case size >= BYTE:
		unit = "B"
	case size == 0:
		return "0"
	}
	stringValue := fmt.Sprintf("%.2f", value)
	stringValue = strings.TrimSuffix(stringValue, ".00")
	return fmt.Sprintf("%s%s", stringValue, unit)
}
