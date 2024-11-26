package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	topic := "notification"

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:19092",
		"client.id":         "myProducer",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer %s\n", err)
		os.Exit(1)
	}

	go func() {
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
				fmt.Printf("consumed message from the queue%s\n", string(e.Value))
			case kafka.Error:
				fmt.Printf("%v\n", e)
			}
		}
	}()

	delivery_chan := make(chan kafka.Event, 10000)

	for {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte("FOO"),
		},
			delivery_chan, // delivery channel
		)
		if err != nil {
			log.Fatal(err)
		}
		_ = <-delivery_chan
		time.Sleep(time.Second * 3)
	}

	//fmt.Printf("%+v\n", e.String())
	//fmt.Printf("%+v\n", p)
}
