package logger

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// Singleton Pattern
func GetLogger() *logrus.Logger {
	var once sync.Once
	if logger == nil {
		once.Do(
			func() {
				logger = logrus.New()

				src, err := os.OpenFile("payment.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

				if err != nil {
					fmt.Print(err.Error())
					fmt.Print("unable to create user.log file")
				}

				multiWriter := io.MultiWriter(os.Stdout, src)

				logger.SetFormatter(&logrus.JSONFormatter{})
				logger.SetOutput(multiWriter)
			})
	}

	return logger
}
