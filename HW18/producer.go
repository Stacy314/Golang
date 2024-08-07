package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"

    "github.com/segmentio/kafka-go"
)

var wg sync.WaitGroup

func Producer(id int, ch chan<- int, writer *kafka.Writer) {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    defer wg.Done()

    for {
        select {
        case <-ticker.C:
            num := rand.Intn(100)
            fmt.Printf("Producer %d produced %d\n", id, num)
            ch <- num

            msg := kafka.Message{
                Key:   []byte(fmt.Sprintf("Key-%d", id)),
                Value: []byte(fmt.Sprintf("%d", num)),
            }
            err := writer.WriteMessages(nil, msg)
            if err != nil {
                fmt.Printf("could not write message to kafka: %v\n", err)
            }
        }
    }
}

func main() {
    ch := make(chan int)

    kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  []string{"localhost:9092"},
        Topic:    "topic-example",
        Balancer: &kafka.LeastBytes{},
    })
    defer kafkaWriter.Close()

    wg.Add(1)
    go Producer(1, ch, kafkaWriter)

    go func() {
        Consumer(1, ch)
    }()

    wg.Wait()
}
