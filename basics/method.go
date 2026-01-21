package basics

import "errors"

func Method() {
	/*Structure of function ( return type is optional nad we can have multiple return types using comma(,) separator)
	func <name> (parameter_list) (return_type){
		block scope of code
	return <value>
	}*/

	result := add(5, 10)
	println("Sum is:", result)

	//Anonymous function
	func(msg string) {
		println(msg)
	}("Hello from Anonymous Function")

	//OR store in variable
	anonFunc := func(x, y int) int {
		return x * y
	}
	product := anonFunc(4, 5)
	println("Product is:", product)

	sum := operate(7, 3, add)
	println("Sum using function as parameter:", sum)

	//return function so need to store in variable to print
	doubleFunc := getMultiplier(2)
	doubleValue := doubleFunc(15)
	println("Double Value:", doubleValue)

	//Variadic Function where we can pass multiple values of same type using (...)
	msg, total := sumAll("Total Sum is 1 + 2 + 3 + 4:", 1, 2, 3, 4)
	println(msg, total)

	res, err := checkValue(10, 10)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	println(res)

	//Directly calling returned function
	// tripleValue := getMultiplier(3)(20)
	// println("Triple Value:", tripleValue)

	//Note: Go supports closures, so the returned function can access variables from the outer function's scope.
}

func add(a, b int) int {
	return a + b
}

// Function as first class citizen (Passing function as parameter)
func operate(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

// Function that return function
func getMultiplier(factor int) func(int) int {
	return func(x int) int {
		print("Geting X value ", x)
		print("Getting factor value ", factor)
		return x * factor
	}
}

// Return multiple type value
func checkValue(a, b int) (string, error) {
	if a > b {
		return "a is greater", nil
	} else if a < b {
		return "b is greater", nil
	}
	return "", errors.New("a and b are equal")
}

//Variadic Function
func sumAll(str string, numbers ...int) (string, int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return str, total
}
