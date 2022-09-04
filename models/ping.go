package models

import "time"

type PingBody struct {
	Ip      string        `json:"ip"`
	Port    int           `json:"port"`
	TimeOut time.Duration `json:"time_out"`
}

type PingStatus struct {
	Status bool `json:"status"`
}
