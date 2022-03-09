package storage

import (
    "log"
    "os"
)

    

func AddLog(StringToAdd string) {
	// If the file doesn't exist, create it or append to the file
    file, err := os.OpenFile("C:\\Users\\Acer\\Desktop\\go-workspace\\booking_app\\storage\\logs\\booking.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)

    log.Println(StringToAdd)
}