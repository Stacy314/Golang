package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"

    "github.com/segmentio/kafka-go"
)

func Consumer(id int, ch <-chan int) {
    defer wg.Done()

    kafkaReader := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   []string{"localhost:9092"},
        Topic:     "topic-example",
        Partition: 0,
        MinBytes:  10e3,
        MaxBytes:  10e6,
    })
    defer kafkaReader.Close()

    for {
        msg, err := kafkaReader.ReadMessage(context.Background())
        if err != nil {
            fmt.Printf("could not read message from kafka: %v\n", err)
            continue
        }

        fmt.Printf("Consumer %d received message: %s = %s\n", id, string(msg.Key), string(msg.Value))
    }
}

func main() {
    ch := make(chan int)

    wg.Add(1)
    go func() {
        Consumer(1, ch)
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-c
        fmt.Println("Received signal, shutting down")
        close(ch)
        wg.Wait()
    }()

    wg.Wait()
}
