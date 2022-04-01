package service

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// type LoggerService struct {
// 	FilePath string
// }

// func NewLogger(moduleName string) *LoggerService {
// 	path := filepath.Join(moduleName, "_logs.log")
// 	return &LoggerService{
// 		FilePath: path,
// 	}
// }

// func (l LoggerService) Log(message string) {
// 	currDate := time.Now().Format("2006-01-02")
// 	filename := l.FilePath + string(os.PathSeparator) + "log_" + currDate + ".log"

// 	message = time.Now().String() + " : " + message + "\n"

// 	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()

// 	if _, err = f.WriteString(message); err != nil {
// 		panic(err)
// 	}
// }

func LogingSetup() {
	var filename string = "logfile.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}
}
