package streamlog

import (
	"fmt"
	"github.com/NubeIO/lib-journalctl/journalctl"
	"github.com/google/uuid"
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

func CreateStreamLog(body *Log) string {
	body.UUID = fmt.Sprintf("log_%s", uuid.New().String())
	body.Message = []string{}
	go createLogStream(body)
	return body.UUID
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
	time.Sleep(time.Duration(body.Duration) * time.Second)
	entries, err := journalctl.NewJournalCTL().EntriesAfter(body.Service, "", "")
	for _, entry := range entries {
		body.Message = append(body.Message, entry.Message)
	}
	if err == nil {
		Logs = append(Logs, body)
	}
}
