package basics

import "fmt"

func Defer() {
	/**Go routines are functions which run in the background, which are running concurrently in background and they are not part of main thread.Not blocking the main thread but in the  background, and then comes back to join the main thread once its job is done.

	the deferred function will be executed after the surrounding function returns, the deferred statment is followed by a function or method call.
	the function is evaluated immediately, but its execution is deferred until the surrounding function returns.so defer function is always part of the another function

	the surrounding function menas the function that encloses the defer statement.
	**/
	process(2)
}


func process(i int) {
	defer fmt.Println("Cleanup")

	fmt.Println("Start")

	if i % 2 == 0 {
		fmt.Println("Processing even number:", i)
		return
	}

	// fmt.Println("Cleanup") if we dont use defer then we have to call cleanup manually before every return statement
}
