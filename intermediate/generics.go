package intermediate

import "fmt"

type Stack[T any] struct {
	elements []T
}

func Generics() {
	/**Generics allow you to write flexible and reusable code by defining functions, types, or data structures that can operate on different data types while maintaining type safety.**/

	x, y := 10, 20
	a, b := "Hello", "World"
	fmt.Println("Before Swap:", x, y)
	x, y = swap(x, y)
	fmt.Println("After Swap:", x, y)
	fmt.Println("Before Swap:", a, b)
	a, b = swap(a, b)
	fmt.Println("After Swap:", a, b)

	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(-23384)
	intStack.Push(9248)
	fmt.Println("Integer Stack:", intStack)

	val, exit := intStack.Pop()
	if exit {
		fmt.Println("Popped from Integer Stack:", val)
	} else {
		fmt.Println("Integer Stack is empty")
	}

	fmt.Println("intStack is Empty: ", intStack.IsEmpty())

	intStack.Print()
}

func swap[T any](a, b T) (T, T) {
	return b, a
}

func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1] //remove last element using slicing 0 - len-1
	return value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) Print() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	for _, val := range s.elements {
		fmt.Print(val)
	}
	fmt.Println()
}
