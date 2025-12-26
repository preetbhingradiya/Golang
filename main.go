package main

import (
	"basicsgo/basics"
	"fmt"
)

func main() {
	fmt.Println("Hello, Golang!")

	x := 43

	fmt.Println("The value of x is:", x)
	basics.PackageExample()
	basics.OverflowUnderflow()
}
