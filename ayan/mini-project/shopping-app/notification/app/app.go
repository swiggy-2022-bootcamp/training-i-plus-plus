package app

import "notification/kafka"

func Start() {
	kafka.ConsumeOrders()
}
