package goKafka

type GoKafkaServices interface {
	WriteMessage(string, interface{}) (bool, error)
}
