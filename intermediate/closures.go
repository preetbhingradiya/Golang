package intermediate

import "fmt"

func Closures() {
	/**closures are a powerful concept in Go that allow functions to capture and manipulate varibles that are defined outside their body.
	A closure is a function value that references variables from outside its body.**/

	sequence := adder() //adder() returns another inner function which is assigned to sequence variable

	//Every time we call sequence(), it increments and returns the next number in the sequence because i is scoped to the adder function.
	//here adder function is called only once but the inner function returned by adder is called multiple times and it retains the state of i between calls.
	fmt.Println("First Call:", sequence())
	fmt.Println("Second Call:", sequence())
	fmt.Println("Third Call:", sequence())

	//Nested Closures
	counterFunc := newCounter()
	counterInner := counterFunc()
	fmt.Println("Counter First Call:", counterInner(30))
}

func adder() func() int {
	i := 0
	fmt.Println("Initial Sum is:", i)
	return func() int {
		i++
		fmt.Println("I value is:", i)
		return i
	}
}

func newCounter() func() func(k int) int {
	j := 0
	fmt.Println("Initial value:", j)

	return func() func(k int) int {
		j++
		fmt.Println("Counter value:", j)

		return func(k int) int {
			j++
			fmt.Println("Counter value in inner function:", j + k)
			return j
		}
	}
}
