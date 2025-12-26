package main

import (
	"basicsgo/basics"
	"fmt"
)

func main() {
	fmt.Println("Hello, Golang!")

	x := 43

	fmt.Println("The value of x is:", x)
	// 1. PACKAGE EXAMPLE
	// basics.PackageExample()

	// 2 Over and Under Flow
	// basics.OverflowUnderflow()

	// 3 Condition and Loop
	// basics.Loop()

	// 4 Guessing Number Game
	basics.GuessingNumber()
}
