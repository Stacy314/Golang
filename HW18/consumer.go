package main

import (
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
)

func Consumer(id int, ch <-chan int) {
    defer wg.Done()

    for num := range ch {
        fmt.Printf("Consumer %d consumed %d\n", id, num)
    }
}

func main() {
    ch := make(chan int)

    wg.Add(2)
    go Producer(1, ch)
    go Consumer(1, ch)

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
