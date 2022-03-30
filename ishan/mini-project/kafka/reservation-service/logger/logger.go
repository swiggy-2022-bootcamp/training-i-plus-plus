package logger

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func Loggerx() *log.Logger {
	year, month, day := time.Now().Date()
	LOG_FILE_LOCATION := fmt.Sprintf("./logs/kafka_%v-%v-%v.log", strconv.Itoa(day), strconv.Itoa(int(month)), strconv.Itoa(year))

	flag.Parse()
	if _, err := os.Stat(LOG_FILE_LOCATION); os.IsNotExist(err) {
		file, err1 := os.Create(LOG_FILE_LOCATION)
		if err1 != nil {
			panic(err1)
		}
		return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		file, err := os.OpenFile(LOG_FILE_LOCATION, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
