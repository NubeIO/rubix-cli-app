package streamlog

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-journalctl/journalctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

var Logs []*Log

type LogType string

type Log struct {
	UUID           string   `json:"uuid"`
	Service        string   `json:"service" binding:"required"`
	Duration       int      `json:"duration" binding:"required"`
	KeyWordsFilter []string `json:"key_words_filter"` // example: mqtt, connected
	Message        []string `json:"message"`
}

func CreateLogAndReturn(body *Log) (*Log, error) {
	streamLog, err := CreateStreamLog(body) // add a log
	if err != nil {
		return nil, err
	}
	timeDuration := body.Duration + 1
	time.Sleep(time.Duration(timeDuration) * time.Second)
	streamLogData := GetStreamLog(streamLog) // get add
	DeleteStreamLog(streamLog)               // delete the log
	return streamLogData, nil                // return the data
}

func GetStreamsLogs() []*Log {
	if Logs == nil {
		Logs = []*Log{}
	}
	return Logs
}

func GetStreamLog(uuid string) *Log {
	for _, _log := range Logs {
		if _log.UUID == uuid {
			return _log
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

func checkSubstrings(str string, subs ...string) (bool, int) {
	matches := 0
	isCompleteMatch := true
	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matches += 1
		} else {
			isCompleteMatch = false
		}
	}
	return isCompleteMatch, matches
}

func createLogStream(body *Log) {
	log.Infof("starting log stream for service: %s for time: %d secounds", body.Service, body.Duration)
	entries, err := journalctl.NewJournalCTL().EntriesAfter(body.Service, "", "")
	for _, entry := range entries {
		lenKeyWordsFilter := len(body.KeyWordsFilter)
		if lenKeyWordsFilter > 0 {
			_, matches := checkSubstrings(entry.Message, body.KeyWordsFilter...)
			if matches == lenKeyWordsFilter {
				body.Message = append(body.Message, entry.Message)
			}
		} else {
			body.Message = append(body.Message, entry.Message)
		}
	}
	if err == nil {
		Logs = append(Logs, body)
	}
	time.Sleep(time.Duration(body.Duration) * time.Second)
	log.Infof("finished log stream for service: %s for time: %d secounds", body.Service, body.Duration)
}
