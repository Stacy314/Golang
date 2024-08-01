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

type Basket struct {
	small  int
	medium int
	large  int
	mu     sync.Mutex
}

func (b *Basket) addOrange(size int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if size <= 5 {
		b.small++
	} else if size <= 10 {
		b.medium++
	} else {
		b.large++
	}
}

func (b *Basket) printCounts() {
	b.mu.Lock()
	defer b.mu.Unlock()

	fmt.Printf("Oranges: small=%d, medium=%d, large=%d\n", b.small, b.medium, b.large)
}

func consumeOranges(brokers []string, topic string, basket *Basket) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var orange Orange
			err := json.Unmarshal(msg.Value, &orange)
			if err != nil {
				fmt.Println("Failed to unmarshal orange:", err)
				continue
			}
			basket.addOrange(orange.Size)
		case <-signals:
			return
		}
	}
}

func main() {

}
