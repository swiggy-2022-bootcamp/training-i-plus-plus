package logger

import (
	"os"
	"time"
)

type Logger struct {
	Filepath string
}

func (l *Logger) Log(message string) {

	currDate := time.Now().Format("2006-01-02")
	filename := l.Filepath + string(os.PathSeparator) + "log_" + currDate + ".log"

	message = time.Now().String() + " : " + message + "\n"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(message); err != nil {
		panic(err)
	}

}
