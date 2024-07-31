package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "time"

    "github.com/IBM/sarama"
)

type Orange struct {
	Size int `json:"size"`
}

func produceOranges(brokers []string, topic string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	for {
		orange := Orange{
			Size: rand.Intn(20) + 1,
		}
		message, err := json.Marshal(orange)
		if err != nil {
			fmt.Println("Failed to marshal orange:", err)
			continue
		}

		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(message),
		}

		_, _, err = producer.SendMessage(msg)
		if err != nil {
			fmt.Println("Failed to send message:", err)
		} else {
			fmt.Println("Sent orange:", orange)
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	brokers := []string{"localhost:9092"}
	topic := "oranges"

	rand.Seed(time.Now().UnixNano())
	produceOranges(brokers, topic)
}
