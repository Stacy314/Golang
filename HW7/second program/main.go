/*Write a program that creates 2 goroutines. The first goroutine generates random 
numbers in a given range and sends them through the channel to the second goroutine. 
The second goroutine receives the random numbers and finds the largest and smallest 
number, then sends them back to the first goroutine through the channel. The first 
goroutine displays the largest and smallest numbers on the screen.*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(min, max, count int, outCh chan int, resultCh chan [2]int, doneCh chan bool) {
	for i := 0; i < count; i++ {
		num := rand.Intn(max-min+1) + min
		outCh <- num
		time.Sleep(1 * time.Second)

		results := <-resultCh
		fmt.Printf("Min: %d, Max: %d\n", results[0], results[1])
	}
	close(outCh)
	doneCh <- true
}

func findMinMax(inCh chan int, outCh chan [2]int, doneCh chan bool) {
	var min, max int
	first := true
	for num := range inCh {
		if first {
			min, max = num, num
			first = false
		} else {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		outCh <- [2]int{min, max}
	}
	close(outCh)
	doneCh <- true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numCount := 10
	numCh := make(chan int)
	resultCh := make(chan [2]int)
	doneCh := make(chan bool, 2)

	go generateNumbers(1, 100, numCount, numCh, resultCh, doneCh)
	go findMinMax(numCh, resultCh, doneCh)

	// Очікуємо завершення всіх горутин
	for i := 0; i < 2; i++ {
		<-doneCh
	}
}