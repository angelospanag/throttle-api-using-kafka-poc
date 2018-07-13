package main

import (
	"net/http"
	"time"

	"github.com/angelospanag/throttle-api-using-kafka-poc/consumer"
	appcontext "github.com/angelospanag/throttle-api-using-kafka-poc/context"
	"github.com/angelospanag/throttle-api-using-kafka-poc/handlers"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
)

func main() {

	// Throttled Kafka topic name - store it in App Context
	// TODO: make the topic name a variable coming from a .toml file
	appcontext.AppContext.TopicName = "throttled_topic"

	// Kafka Producer - store it in App Context
	// TODO: make the kafka instance URL a variable coming from a .toml file
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	appcontext.AppContext.KafkaProducer = p

	// Kafka Consumer - store it in App Context
	// TODO: make the kafka instance URL a variable coming from a .toml file
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{appcontext.AppContext.TopicName}, nil)
	appcontext.AppContext.KafkaConsumer = c

	// Ticker for consuming from Kafka every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				consumer.ConsumeFromKafka()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ProduceToKafka).Methods("POST")
	http.ListenAndServe(":8080", r)
}
