package consumer

import (
	"fmt"

	appcontext "github.com/angelospanag/throttle-api-using-kafka-poc/context"
)

// ConsumeFromKafka consumes a specific number of messages from Kafka
// TODO: make the number of messages value coming from an .env file
func ConsumeFromKafka() {

	//appcontext.AppContext.KafkaConsumer.Seek(*appcontext.AppContext.TopicPartition, 0)

	for i := 0; i <= 4; i++ {
		msg, err := appcontext.AppContext.KafkaConsumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			appcontext.AppContext.TopicPartition = &msg.TopicPartition
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}
