package log

import (
	"log"
	"os"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	fileMode      = false
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

	if fileMode {
		InfoLogger.SetOutput(file)
		WarningLogger.SetOutput(file)
		ErrorLogger.SetOutput(file)
	}
}

func Info(data ...interface{}) {
	InfoLogger.Print(data, "\n")
}

func Warn(data ...interface{}) {
	WarningLogger.Print(data, "\n")
}

func Error(data ...interface{}) {
	ErrorLogger.Print(data, "\n")
}
