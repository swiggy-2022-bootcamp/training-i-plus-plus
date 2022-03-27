package goKafka

type GoKafkaServices interface {
	StoreOrders(string) error
}
