package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(min, max int, outCh chan int, resultCh chan [2]int) {
	for {
		num := rand.Intn(max-min+1) + min
		outCh <- num
		time.Sleep(1 * time.Second)

		results := <-resultCh
		fmt.Printf("Min: %d, Max: %d\n", results[0], results[1])
	}
}

func findMinMax(inCh chan int, outCh chan [2]int) {
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
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numCh := make(chan int)
	resultCh := make(chan [2]int)

	go generateNumbers(1, 100, numCh, resultCh)
	go findMinMax(numCh, resultCh)

	select {}
}