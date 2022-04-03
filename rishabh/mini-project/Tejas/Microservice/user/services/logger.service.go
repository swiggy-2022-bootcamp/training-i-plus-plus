package services

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Logger interface {
	Log(values ...interface{})
}

type LoggerService struct {
	filename string
	topic    string
}

func (l *LoggerService) Log(values ...interface{}) {
	time := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintln(values...)
	msg = time + " [" + l.topic + "] " + msg

	file, err := os.OpenFile(l.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Mkdir("logs", 0755)
		file, err = os.OpenFile(l.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer file.Close()
		file.WriteString(msg)
		if err != nil {
			panic(err)
		}
	} else {
		defer file.Close()
		file.WriteString(msg)
	}
}

func NewLoggerService(topic string) Logger {
	time := time.Now().Format("2006-01-02")
	path := filepath.Join("logs", time+".log")
	return &LoggerService{
		filename: path,
		topic:    topic,
	}
}
