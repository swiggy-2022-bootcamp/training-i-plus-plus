package app

import (
	"context"
	"notification/kafka"
)

func Start() {

	kafka.ConsumeOrders(context.Background())
}
