package basics

import "fmt"

func OverflowUnderflow() {

	//Overflow with signed integers ( where number is pointer to negative range)
	var a int8 = 127
	fmt.Println("Initial value of a (int8):", a)
	a = a + 1
	fmt.Println("Value of a after overflow (int8):", a) // Should wrap around to -128

	//Overflow with unsigned integers
	var b uint8 = 255
	fmt.Println("Initial value of b (uint8):", b)
	b = b + 1
	fmt.Println("Value of b after overflow (uint8):", b) // Should wrap around to 0

	//Underflow with signed integers
	var c int8 = -128
	fmt.Println("Initial value of c (int8):", c)
	c = c - 1
	fmt.Println("Value of c after underflow (int8):", c) // Should wrap around to 127

	//Underflow with unsigned integers
	var d uint8 = 0
	fmt.Println("Initial value of d (uint8):", d)
	d = d - 1
	fmt.Println("Value of d after underflow (uint8):", d) // Should wrap around to 255

}
