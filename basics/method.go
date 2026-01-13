package basics

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
