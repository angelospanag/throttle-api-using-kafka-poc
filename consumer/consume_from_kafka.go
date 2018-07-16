package consumer

import (
	"fmt"

	appconfig "github.com/angelospanag/throttle-api-using-kafka-poc/config"
	appcontext "github.com/angelospanag/throttle-api-using-kafka-poc/context"
)

// ConsumeFromKafka consumes a specific number of messages from Kafka
// TODO: offsets so that the Kafka Consumer knows where to start processing stored requests if the server ever stops/crashes
func ConsumeFromKafka() {

	//appcontext.AppContext.KafkaConsumer.Seek(*appcontext.AppContext.TopicPartition, 0)

	for i := 1; i <= appconfig.AppConfig.ConsumptionConfig.NumberOfMessages; i++ {
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
