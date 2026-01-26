package basics

import (
	"fmt"
	"log"
)

func Panic() {
	// panic is used to stop normal program execution immediately when the program reaches a state it cannot safely continue from.
	//recover() only works when called directly inside a deferred function.
	panicWithDefer(8)
}

func panicWithDefer(n int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			log.Printf("panic recovered: %v", r)
		}
	}()

	if n > 5 {
		panic("n is greater than 5")
		//Anthing after panic will not be executed at runtime
	}
	fmt.Println("n is:", n)
}
