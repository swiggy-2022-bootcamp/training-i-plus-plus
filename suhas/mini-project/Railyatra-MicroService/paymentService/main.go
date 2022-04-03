package main

import (
	"fmt"
	"paymentService/config"
	"paymentService/controllers"
	log "paymentService/logger"
)

var (
	errLog = log.InfoLogger.Println
)

func main() {
	config.ConnectDB()
	err := controllers.GrpcPaymentServer()
	if err != nil {
		fmt.Println(err)
		return
	}
}
