package streamlog

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-journalctl/journalctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

var Logs []*Log

type Log struct {
	UUID     string   `json:"uuid"`
	Service  string   `json:"service" binding:"required"`
	Duration int      `json:"duration" binding:"required"`
	Message  []string `json:"message"`
}

func GetStreamsLogs() []*Log {
	if Logs == nil {
		Logs = []*Log{}
	}
	return Logs
}

func GetStreamLog(uuid string) *Log {
	for _, log := range Logs {
		if log.UUID == uuid {
			return log
		}
	}
	return nil
}

func CreateStreamLog(body *Log) (string, error) {
	body.UUID = fmt.Sprintf("log_%s", uuid.New().String())
	body.Message = []string{}
	s := systemctl.New(false, 30)
	isRunning, status, err := s.IsRunning(body.Service)
	if !isRunning || err != nil {
		return status, errors.New(fmt.Sprintf("service not running %s", body.Service))
	}
	go createLogStream(body)
	return body.UUID, nil
}

func DeleteStreamLog(uuid string) bool {
	deleted := false
	for i, entry := range Logs {
		if entry.UUID == uuid {
			Logs = append(Logs[:i], Logs[i+1:]...)
			deleted = true
			break
		}
	}
	return deleted
}

func DeleteStreamLogs() {
	Logs = []*Log{}
}

func createLogStream(body *Log) {
	log.Infof("start log stream for service: %s for time: %d secounds", body.Service, body.Duration)
	entries, err := journalctl.NewJournalCTL().EntriesAfter(body.Service, "", "")
	for _, entry := range entries {
		body.Message = append(body.Message, entry.Message)
	}
	if err == nil {
		Logs = append(Logs, body)
	}
	time.Sleep(time.Duration(body.Duration) * time.Second)
	log.Infof("finshed log stream for service: %s for time: %d secounds", body.Service, body.Duration)
}
