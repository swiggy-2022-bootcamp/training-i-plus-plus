package logger

import (
	"log"
	"os"
)

var (
	//InfoLogger ..
	InfoLogger *log.Logger
	//WarningLogger ..
	WarningLogger *log.Logger
	//ErrorLogger ..
	ErrorLogger *log.Logger
	//fileMode ..
	fileMode = true
)

func init() {
	file, err := os.OpenFile("./logs/medoAppLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(os.Stdout, "❕INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(os.Stdout, "⚠️WARNING: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stdout, "📛ERROR: ", log.Ldate|log.Ltime)

	if fileMode {
		InfoLogger.SetOutput(file)
		WarningLogger.SetOutput(file)
		ErrorLogger.SetOutput(file)
	}
}

//Info ..
func Info(data ...interface{}) {
	InfoLogger.Print(data, "\n")
}

//Warn ..
func Warn(data ...interface{}) {
	WarningLogger.Print(data, "\n")
}

//Error ..
func Error(data ...interface{}) {
	ErrorLogger.Print(data, "\n")
}
