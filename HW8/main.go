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
)

type Question struct {
	Question string
	Options  []string
	Answer   int
}

type Player struct {
	Name    string
	Answer  chan int
	Results chan Result
}

type Result struct {
	PlayerName string
	Correct    bool
}

func questionGenerator(ctx context.Context, wg *sync.WaitGroup, questions chan<- Question) {
	defer wg.Done()
	questionsList := []Question{
		{"What is the capital of France?", []string{"Berlin", "Madrid", "Paris", "Rome"}, 2},
		{"What is 2 + 2?", []string{"3", "4", "5", "6"}, 1},
		{"What is the color of the sky?", []string{"Blue", "Green", "Red", "Yellow"}, 0},
	}
	rand.Seed(time.Now().UnixNano())

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Question generator shutting down...")
			return
		case <-time.After(10 * time.Second):
			question := questionsList[rand.Intn(len(questionsList))]
			fmt.Println("New question generated:", question.Question)
			questions <- question
		}
	}
}

func player(ctx context.Context, wg *sync.WaitGroup, player Player, questions <-chan Question) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Player %s shutting down...\n", player.Name)
			return
		case question := <-questions:
			fmt.Printf("Player %s received question: %s\n", player.Name, question.Question)
			answer := rand.Intn(len(question.Options))
			player.Answer <- answer
		}
	}
}

func checker(ctx context.Context, wg *sync.WaitGroup, players []Player, questions <-chan Question, results chan<- Result) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Checker shutting down...")
			return
		case question := <-questions:
			for _, player := range players {
				select {
				case <-ctx.Done():
					fmt.Println("Checker shutting down...")
					return
				case answer := <-player.Answer:
					correct := answer == question.Answer
					results <- Result{PlayerName: player.Name, Correct: correct}
				}
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	questions := make(chan Question)
	results := make(chan Result)

	players := []Player{
		{Name: "Alice", Answer: make(chan int), Results: results},
		{Name: "Bob", Answer: make(chan int), Results: results},
		{Name: "Charlie", Answer: make(chan int), Results: results},
	}

	wg.Add(1)
	go questionGenerator(ctx, wg, questions)

	for _, player := range players {
		wg.Add(1)
		go player(ctx, wg, player, questions)
	}

	wg.Add(1)
	go checker(ctx, wg, players, questions, results)

	go func() {
		for result := range results {
			fmt.Printf("Player %s answered correctly: %t\n", result.PlayerName, result.Correct)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	fmt.Println("Shutting down...")
	cancel()
	wg.Wait()
	fmt.Println("Program terminated.")
}