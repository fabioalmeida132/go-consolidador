package consumer

import (
	"context"
	"fmt"
	"github.com/fabioalmeida132/go-consolidador/internal/infra/kafka/factory"
	"github.com/fabioalmeida132/go-consolidador/pkg/uow"
)

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "gostats",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	kafkaConsumer.SubscribeTopics(topics, nil)

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}

func ProcessEvents(ctx context.Context, msgsChan chan *kafka.Message, uow uow.UowInterface) {

	for msg := range msgsChan {
		fmt.Printf("Received message", string(msg.Value), "on topic", msg.TopicPartition.Topic)
		strategy := factory.CreateProcessMessageStrategy(*msg.TopicPartition.Topic)
		err := strategy.Process(ctx, msg, uow)
		if err != nil {
			fmt.Printf("Error processing message", string(msg.Value), "on topic", msg.TopicPartition.Topic)
		}
	}

}
