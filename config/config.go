package config

import (
	"errors"

	"github.com/spf13/viper"
)

// AppConfig holds the global application configuration that can be used in different packages.
var AppConfig Config

// Config contains all the configuration details of the application.
type Config struct {
	KafkaConfig       *KafkaConfig
	ConsumptionConfig *ConsumptionConfig
}

// KafkaConfig contains all the configuration details of Kafka instances and Topics.
type KafkaConfig struct {
	Servers []string
	Topic   string
}

// ConsumptionConfig contains all the configuration details of for message consumption
// from a Kafka Topic.
type ConsumptionConfig struct {
	NumberOfMessages  int
	TimePeriodSeconds int
}

// InitiateConfig initiates the application configuration.
func InitiateConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("No configuration file loaded")
	}

	// Kafka configuration from configuration file.
	kafkaServers := viper.GetStringSlice("kafka.servers")
	kafkaTopic := viper.GetString("kafka.topic")

	if len(kafkaServers) <= 0 {
		return errors.New("Missing Apache Kafka server URLs")
	}

	if len(kafkaTopic) <= 0 {
		return errors.New("Missing Apache Kafka Topic")
	}

	kafkaConfig := KafkaConfig{
		Servers: kafkaServers,
		Topic:   kafkaTopic,
	}

	AppConfig.KafkaConfig = &kafkaConfig

	// Consumption configuration from configuration file.
	numberOfMessages := 0
	timePeriodSeconds := 0
	numberOfMessages = viper.GetInt("consumption.number_of_messages")
	timePeriodSeconds = viper.GetInt("consumption.time_period_seconds")

	if numberOfMessages <= 0 {
		return errors.New("Number of messages for Consumption not set")
	}

	if timePeriodSeconds <= 0 {
		return errors.New("Time Period in Seconds for Consumption not set")
	}

	consumptionConfig := ConsumptionConfig{
		NumberOfMessages:  numberOfMessages,
		TimePeriodSeconds: timePeriodSeconds,
	}

	AppConfig.ConsumptionConfig = &consumptionConfig

	return nil
}
