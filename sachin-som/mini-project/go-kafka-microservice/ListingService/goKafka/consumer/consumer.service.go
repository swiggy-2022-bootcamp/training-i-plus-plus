package goKafka

type GoKafkaServices interface {
	ReadMessage(string) (interface{}, error)
}
