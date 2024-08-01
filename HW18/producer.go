package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

type Orange struct {
	Size int `json:"size"`
}

func produceOranges(topic string, brokerAddress string) {
	writer := kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	for {
		orange := Orange{Size: rand.Intn(20) + 1}
		message, err := json.Marshal(orange)
		if err != nil {
			fmt.Printf("Error marshaling orange: %v\n", err)
			continue
		}

		err = writer.WriteMessages(
			context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("key-%d", orange.Size)),
				Value: message,
			},
		)
		if err != nil {
			fmt.Printf("Error writing message to kafka: %v\n", err)
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	topic := "oranges"
	brokerAddress := "localhost:9092"

	produceOranges(topic, brokerAddress)
}
