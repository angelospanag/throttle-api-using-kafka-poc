package context

import "github.com/confluentinc/confluent-kafka-go/kafka"

// AppContext holds the global application context that can be used in different packages.
var AppContext Env

// Env defines the stucture of the application context.
type Env struct {
	KafkaProducer  *kafka.Producer
	KafkaConsumer  *kafka.Consumer
	TopicPartition *kafka.TopicPartition
}
