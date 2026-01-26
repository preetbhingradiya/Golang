package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func GuessingNumber() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	targetNumber := random.Intn(100) + 1 // Random number between 1 and 100

	println("Welcome to the Guessing Number Game!")
	println("I have selected a number between 1 and 100. Can you guess it?")

	var guessedNumber int
	for {
		print("Enter your guess: ")
		_, err := fmt.Scan(&guessedNumber)
		if err != nil {
			println("Please enter a valid number.")
			continue
		}
		if guessedNumber < targetNumber {
			println("Too low! Try again.")
		} else if guessedNumber > targetNumber {
			println("Too high! Try again.")
		} else {
			println("Congratulations! You've guessed the number:", targetNumber)
			break
		}
	}

}
