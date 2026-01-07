package basics

import "fmt"

func Array() {
	// Array is a collection of elements of same type
	// 1 Declaration of an array
	var arr [5]int
	fmt.Println("Array:", arr)

	//2 Assigning values to array Declare + init values
	var arr2 [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array 2:", arr2)
	//3 short declaration
	language := [3]string{"Golang", "Python", "Java"}
	fmt.Println("Array 3:", language)

	code := language
	fmt.Println("Code:", code)

	//4 auto size array
	arr4 := [...]float64{3.14, 1.618, 2.718}
	fmt.Println("Array 4:", arr4)

	//5 index based access
	arr5 := [6]int{1: 10, 4: 50}
	fmt.Println("Array 5:", arr5)
	// Accessing elements
	fmt.Println("Element at index 1 of arr5:", arr5[1])

	//6 Modifying elements
	arr5[2] = 20
	fmt.Println("Modified Array 5:", arr5)

	//pointer and address of array
	//here copyArr is a pointer to an array of 3 integers and this holds address of orginalArr
	originalArr := [3]int{1, 2, 3}
	var copyArr *[3]int
	copyArr = &originalArr

	copyArr[1] = 999 //modifying value using pointer
	fmt.Println("Original Array:", originalArr)
	fmt.Println("Copy Array (Pointer):", *copyArr) // this not varible use as pointer

}
