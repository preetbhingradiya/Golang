package basics

import (
	"fmt"
	"slices"
)

func Slice() {
	// Slice is a dynamically sized, flexible view into the elements of an array.
	// 1 Declaration of a slice
	var s []int
	fmt.Println("Slice:", s)

	a := [5]int{10, 20, 30, 40, 50}
	s = a[1:4] //slicing array a from index 1 to 3 where index 4 is excluded (ex. [:] means full array)
	fmt.Println("Sliced Array (Slice):", s)

	//also use make function to create slice but here we need to provide length like array
	slice2 := make([]string, 3, 5) //length 3
	slice2[0] = "Golang"
	slice2[1] = "Python"
	slice2[2] = "Java"
	slice2 = append(slice2, "C++") //adding new element to slice
	fmt.Println("Slice 2:", slice2)
	fmt.Println("Length of Slice2: ", len(slice2), "\nCapacity of Slice2 :", cap(slice2))

	//Passing array using slice
	arr := [3]int{1, 2, 3}
	update(arr[:]) //passing array as slice
	fmt.Println("Array after update function:", arr)

	//Comparing slices
	if slices.Equal(s, []int{20, 30, 40}) {
		fmt.Println("Slices are equal")
	}
}

func update(s []int) {
	s[0] = 999
}
