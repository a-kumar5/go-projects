package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	topic := "notification"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:19092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest",
	})

	if err != nil {
		log.Fatal(err)
	}
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("consumed message from the queue %s\n", string(e.Value))
		case kafka.Error:
			fmt.Printf("%v\n", e)
		}
	}
}
