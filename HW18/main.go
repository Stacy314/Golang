package main

import (
	"sync"
	"HW18/producer"
	"HW18/consumer"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		producer.RunProducer()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		consumer.RunConsumer()
	}()

	wg.Wait()
}
