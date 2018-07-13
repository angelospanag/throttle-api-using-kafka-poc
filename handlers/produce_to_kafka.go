package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	appcontext "github.com/angelospanag/throttle-api-using-kafka-poc/context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// ProduceToKafka is an endpoint that returns an HTTP OK status a
func ProduceToKafka(w http.ResponseWriter, r *http.Request) {

	var err error
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}
	defer r.Body.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range appcontext.AppContext.KafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	go func() {

		appcontext.AppContext.KafkaProducer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &appcontext.AppContext.TopicName, Partition: kafka.PartitionAny},
			Value:          b,
		}, nil)
	}()

	// Return a JSON response to the user with status OK
	w.Write([]byte(`{"status":"OK"}`))
}
