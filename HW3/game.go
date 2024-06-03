package main

import (
	"fmt"
	"os"
)

type Game struct {
	name       string
	inventory  []string
	location   string
	isAlive    bool
	isBitten   bool
}

func main() {
	game := Game{
		location: "cave",
		isAlive:  true,
	}
	fmt.Printf("You wake up at the entrance of a cave. You don't remember anything. \n")
	fmt.Println("Enter your name ")
	fmt.Scanf("%s", &game.name)
	fmt.Printf("The only thing you remember is your name - %s. There is a backpack next to you.\n", game.name)
	game.inventory = []string{"matches", "flashlight", "knife"}

	game.play()
}

func (game *Game) play() {
	for game.isAlive {
		switch game.location {
		case "cave":
			game.cave()
		case "forest":
			game.forest()
		case "camp":
			game.camp()
		case "safe":
			game.safe()
		default:
			game.isAlive = false
		}
	}

	if game.isBitten {
		fmt.Println("You have been bitten and fainted. Game over.")
	} else {
		fmt.Println("Thank you for playing.")
	}
}

func (game *Game) cave() {
	var choice string
	fmt.Println("It's dark in the cave. Do you want to explore the cave or leave for the forest? (explore/leave)")
	fmt.Scanf("%s", &choice)

	switch choice {
	case "explore":
		fmt.Println("You decided to explore the cave. Unfortunately, it's too dark and you fall into a pit. Game over.")
		game.isAlive = false
	case "leave":
		fmt.Println("You decided to leave the cave and head to the forest.")
		game.location = "forest"
	default:
		fmt.Println("Invalid choice. Try again.")
	}
}

func (game *Game) forest() {
	var choice string
	fmt.Println("You are in the forest. There is a dead animal nearby. Do you want to investigate or move on? (investigate/move)")
	fmt.Scanf("%s", &choice)

	switch choice {
	case "investigate":
		fmt.Println("You decided to investigate the dead animal. It looks strange but you find nothing useful. You continue walking and reach an empty camp.")
		game.location = "camp"
	case "move":
		fmt.Println("You decided to ignore the animal and move on. You reach an empty camp.")
		game.location = "camp"
	default:
		fmt.Println("Invalid choice. Try again.")
	}
}

func (game *Game) camp() {
	var choice string
	fmt.Println("You are at an empty camp. You are tired. Do you want to rest or keep moving? (rest/move)")
	fmt.Scanf("%s", &choice)

	switch choice {
	case "rest":
		fmt.Println("You decided to rest. In the nearest tent, you find a safe with a two-digit code lock. Do you want to try opening it? (yes/no)")
		fmt.Scanf("%s", &choice)
		if choice == "yes" {
			game.location = "safe"
		} else {
			fmt.Println("You decided not to open the safe and rest. However, without any further action, you stay there indefinitely. Game over.")
			game.isAlive = false
		}
	case "move":
		fmt.Println("You decided to keep moving but soon collapse from exhaustion. Game over.")
		game.isAlive = false
	default:
		fmt.Println("Invalid choice. Try again.")
	}
}

func (game *Game) safe() {
	var code int
	fmt.Println("Enter the two-digit code to open the safe. (Hint: Answer to the Ultimate Question of Life, the Universe, and Everything):")
	fmt.Scanf("%d", &code)

	if code == 42 {
		fmt.Println("The safe opens, you learned all the secrets of the universe.")
		os.Exit(0)
	} else {
		fmt.Println("Incorrect code. The safe does not open. You feel frustrated and decide to rest. Game over.")
		game.isAlive = false
	}
}