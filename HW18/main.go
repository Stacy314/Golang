package main

import (
    "context"
    "fmt"
    "math/rand"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"

    "github.com/segmentio/kafka-go"
)

var wg sync.WaitGroup

func Producer(id int, ch chan<- int, writer *kafka.Writer, done <-chan struct{}) {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    defer wg.Done()

    for {
        select {
        case <-ticker.C:
            num := rand.Intn(100)
            fmt.Printf("Producer %d produced %d\n", id, num)

            select {
            case ch <- num:
                // Send the number to Kafka
                msg := kafka.Message{
                    Key:   []byte(fmt.Sprintf("Key-%d", id)),
                    Value: []byte(fmt.Sprintf("%d", num)),
                }
                err := writer.WriteMessages(context.Background(), msg)
                if err != nil {
                    fmt.Printf("could not write message to kafka: %v\n", err)
                }
            case <-done:
                return
            }
        case <-done:
            return
        }
    }
}

func Consumer(id int, ch <-chan int, done <-chan struct{}) {
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
        select {
        case msg := <-ch:
            fmt.Printf("Consumer %d received message: %d\n", id, msg)
        case <-done:
            return
        }
    }
}

func main() {
    ch := make(chan int)
    done := make(chan struct{})

    kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  []string{"localhost:9092"},
        Topic:    "topic-example",
        Balancer: &kafka.LeastBytes{},
    })
    defer kafkaWriter.Close()

    // Wait until Kafka is ready
    for {
        conn, err := kafka.Dial("tcp", "localhost:9092")
        if err != nil {
            fmt.Printf("Kafka is not ready, retrying in 1 second...\n")
            time.Sleep(1 * time.Second)
        } else {
            conn.Close()
            break
        }
    }

    wg.Add(1)
    go Producer(1, ch, kafkaWriter, done)

    wg.Add(1)
    go Consumer(1, ch, done)

    // Capture interrupt signal for graceful shutdown
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-c
        fmt.Println("Received signal, shutting down")
        close(done)
        wg.Wait()
        close(ch)
        os.Exit(0)
    }()

    wg.Wait()
}
