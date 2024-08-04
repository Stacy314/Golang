package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var wg sync.WaitGroup

func Producer(id int, ch chan<- int) {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    defer wg.Done()

    for {
        select {
        case <-ticker.C:
            num := rand.Intn(100)
            fmt.Printf("Producer %d produced %d\n", id, num)
            ch <- num
        }
    }
}

func main() {
    ch := make(chan int)

    wg.Add(2)
    go Producer(1, ch)
    go Consumer(1, ch)

    wg.Wait()
}

