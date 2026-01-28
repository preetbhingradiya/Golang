package intermediate

import "fmt"

func Pointer() {
	/**Pointers are variables that store the memory address of another variable.**/

	//var ptr *int
	var ptr *int //pointer variable that can hold the address of an integer variable
	var a int = 99
	ptr = &a //assigning the address of variable a to pointer ptr using & operator

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of ptr (Address of a):", ptr)
	fmt.Println("Value at the address stored in ptr:", *ptr) //Dereferencing the pointer to get the value at the address it points to using * operator

	//Modifying value using pointer
	modifyValue(ptr)
	fmt.Println("Value of a after modification through pointer:", a)
}

func modifyValue(val *int) {
	*val = *val + 10 //Dereferencing the pointer to modify the value at the address
}
