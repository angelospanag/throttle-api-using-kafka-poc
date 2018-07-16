package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	appconfig "github.com/angelospanag/throttle-api-using-kafka-poc/config"
	"github.com/angelospanag/throttle-api-using-kafka-poc/consumer"
	appcontext "github.com/angelospanag/throttle-api-using-kafka-poc/context"
	"github.com/angelospanag/throttle-api-using-kafka-poc/handlers"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
)

func main() {

	var err error

	err = appconfig.InitiateConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Kafka Producer - store it in App Context
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": strings.Join(appconfig.AppConfig.KafkaConfig.Servers, ", ")})
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
	c.SubscribeTopics([]string{appconfig.AppConfig.KafkaConfig.Topic}, nil)
	appcontext.AppContext.KafkaConsumer = c

	// Ticker for consuming from Kafka every x seconds
	ticker := time.NewTicker(time.Duration(appconfig.AppConfig.ConsumptionConfig.TimePeriodSeconds) * time.Second)
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
