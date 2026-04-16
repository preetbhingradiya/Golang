package intermediate

import (
	"fmt"
	"math/rand"
)

func DiceGame() {

	fmt.Println("Welcome to the dice game")
	for {
		// Simulate rolling a dice
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")

		var choice int
		_, err := fmt.Scanln(&choice)

		if err != nil || choice < 1 || choice > 2 {
			fmt.Println("Invalid choice. Please enter 1 or 2.")
			continue
		}

		if choice == 2 {
			fmt.Println("Exiting the game. Goodbye!")
			break
		}

		diceRoll := rand.Intn(6) + 1 // Generates a random number between 1 and 6
		fmt.Printf("You rolled a %d\n", diceRoll)

		if diceRoll == 6 {
			fmt.Println("Congratulations! You win!")
			break
		} else {
			fmt.Println("Try again!")
		}
	}

}
