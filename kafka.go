package main

import (

	"github.com/IBM/sarama"
)


func connectToConsumer(brokerUrls []string) (sarama.Consumer, error) {
	config := *sarama.NewConfig()
	config.Producer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokerUrls, &config)
	return consumer, err
}