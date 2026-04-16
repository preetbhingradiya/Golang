package main

import (
	// "basicsgo/basics"

	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, Golang!")
	/*
		// 1. PACKAGE EXAMPLE
		basics.PackageExample()
		// 2 Over and Under Flow
		basics.OverflowUnderflow()
		// 3 Condition and Loop
		basics.Loop()
		// 4 Guessing Number Game
		basics.GuessingNumber()
		//5 Arrays
		 basics.Array()
		//6 Slices
		 basics.Slice()
		//7 Function
		 basics.Method()
		//8 Defer
		 basics.Defer()
		//9 Panic and Recover
		 basics.Panic()
	*/

	//Intermediate Level
	/**
	 //1. Closures
	 intermediate.Closures()
	 //2 Pointers
	 intermediate.Pointer()
	 //3 Structs
	 intermediate.Struct()
	 //4 Interfaces
	 intermediate.Interface()
	 //5 Generics
	 intermediate.Generics()
	 //6 Custom Error
	 intermediate.CustomError()
	 //7 URL Parsing
	 intermediate.UrlParsing()

	 intermediate.DiceGame()
	**/

	// time := time.Now()
	// fmt.Println("Current Time:", time)
	// fmt.Println("Year:", time.Format("2006-01-02"))
	// epoch := time.Unix()
	// fmt.Println("Epoch Time:", epoch)

	fmt.Print(rand.Intn(101)) // Generates a random number between 0 and 100

	val := rand.New(rand.NewSource(42)) //fixed seed for reproducibility
	fmt.Print(val.Intn(101))

}
