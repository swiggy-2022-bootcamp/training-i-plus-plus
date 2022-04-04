package service

import (
	"io"
	"log"
	//"User-Service/server"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

var GetDefaultWriter *io.Writer

func InitLoggerService() {
	InfoLogger = log.New(*GetDefaultWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(*GetDefaultWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(*GetDefaultWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
