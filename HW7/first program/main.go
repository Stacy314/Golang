/*Write a program that creates 3 goroutines. The first goroutine 
generates random numbers and sends them through the channel to the 
second goroutine. The second goroutine takes the random numbers 
and averages them, then sends it to the third goroutine over 
the channel. The third goroutine displays the average value on 
the screen.*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(ch chan int, count int) {
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		ch <- num
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func calculateAverage(inCh chan int, outCh chan float64) {
	var sum int
	var count int
	for num := range inCh {
		sum += num
		count++
		average := float64(sum) / float64(count)
		outCh <- average
	}
	close(outCh)
}

func printAverage(ch chan float64) {
	for avg := range ch {
		fmt.Printf("Average: %f\n", avg)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numCount := 10
	numCh := make(chan int)
	avgCh := make(chan float64)

	go generateNumbers(numCh, numCount)
	go calculateAverage(numCh, avgCh)
	go printAverage(avgCh)

	time.Sleep(time.Duration(numCount+2) * time.Second)
}