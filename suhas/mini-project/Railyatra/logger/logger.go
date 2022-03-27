package logger

import (
	"log"
	"os"
)

func ALog(StringToAdd string) {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(".\\logs\\Admin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println(StringToAdd)
}

func ULog(StringToAdd string) {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(".\\logs\\Admin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println(StringToAdd)
}

func Klog(StringToAdd string) {
	file, err := os.OpenFile(".\\logs\\Kafka.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println(StringToAdd)
}
