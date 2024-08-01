package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type Orange struct {
	Size int `json:"size"`
}

type Basket struct {
	Small  int
	Medium int
	Large  int
}

func consumeOranges(topic string, brokerAddress string, basket *Basket) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "orange-consumers",
	})

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Error reading message from kafka: %v\n", err)
			continue
		}

		var orange Orange
		err = json.Unmarshal(m.Value, &orange)
		if err != nil {
			fmt.Printf("Error unmarshaling message: %v\n", err)
			continue
		}

		if orange.Size < 7 {
			basket.Small++
		} else if orange.Size < 15 {
			basket.Medium++
		} else {
			basket.Large++
		}
	}
}

func printStats(basket *Basket) {
	for {
		time.Sleep(10 * time.Second)
		fmt.Printf("Oranges: small=%d, medium=%d, large=%d\n", basket.Small, basket.Medium, basket.Large)
	}
}

func main() {
	topic := "oranges"
	brokerAddress := "localhost:9092"
	basket := &Basket{}

	go consumeOranges(topic, brokerAddress, basket)
	printStats(basket)
}
