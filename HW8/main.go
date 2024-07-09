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
	defer close(questions) // Close the channel when done
	questionsList := []Question{
		{"What is the capital of France?", []string{"Berlin", "Madrid", "Paris", "Rome"}, 2},
		{"What is 2 + 2?", []string{"3", "4", "5", "6"}, 1},
		{"What is the color of the sky?", []string{"Blue", "Green", "Red", "Yellow"}, 0},
	}
	rand.Seed(time.Now().UnixNano())

	for _, q := range questionsList {
		select {
		case <-ctx.Done():
			fmt.Println("Question generator shutting down...")
			return
		case <-time.After(10 * time.Second):
			fmt.Println("New question generated:", q.Question)
			questions <- q
		}
	}
}

func player(ctx context.Context, wg *sync.WaitGroup, player Player, questions <-chan Question, questionCount int, done chan struct{}) {
	defer wg.Done()
	defer close(player.Answer) // Close the channel when done

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Player %s shutting down...\n", player.Name)
			return
		case question, ok := <-questions:
			if !ok {
				return
			}
			fmt.Printf("Player %s received question: %s\n", player.Name, question.Question)
			answer := rand.Intn(len(question.Options))
			player.Answer <- answer

			// Check if all questions have been answered
			if questionCount--; questionCount == 0 {
				close(done)
				return
			}
		}
	}
}

func checker(ctx context.Context, wg *sync.WaitGroup, players []Player, results chan<- Result) {
	defer wg.Done()
	defer close(results) // Close the channel when done

	for {
		for _, player := range players {
			select {
			case <-ctx.Done():
				fmt.Println("Checker shutting down...")
				return
			case answer, ok := <-player.Answer:
				if !ok {
					return
				}
				correct := answer == rand.Intn(4) // Random check
				results <- Result{PlayerName: player.Name, Correct: correct}
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	questions := make(chan Question)
	results := make(chan Result)
	done := make(chan struct{}) // Signal channel for all questions answered

	players := []Player{
		{Name: "Alice", Answer: make(chan int), Results: results},
		{Name: "Bob", Answer: make(chan int), Results: results},
		{Name: "Charlie", Answer: make(chan int), Results: results},
	}

	wg.Add(1)
	go questionGenerator(ctx, wg, questions)

	questionCount := len(players) * 3 // Assuming each player answers all questions
	for _, p := range players {
		wg.Add(1)
		go player(ctx, wg, p, questions, questionCount, done)
	}

	wg.Add(1)
	go checker(ctx, wg, players, results)

	go func() {
		<-done // Wait until all questions are answered
		cancel()
	}()

	go func() {
		for result := range results {
			fmt.Printf("Player %s answered correctly: %t\n", result.PlayerName, result.Correct)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	fmt.Println("Shutting down...")
	wg.Wait()
	fmt.Println("Program terminated.")
}