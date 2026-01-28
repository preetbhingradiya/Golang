package intermediate

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func Interface() {
	/**Interfaces are abstract types that define a set of method signatures. They allow different types to be treated uniformly based on shared behavior.**/

	r := Rect{width: 10, height: 5}
	c := Circle{radius: 7}

	measure(r)
	measure(c)

	//Print the message
	printMessage("Hello, World!", 42, 3.14, true)
}

func printMessage(i ...interface{}) {
	for _, msg := range i {
		fmt.Println(msg)
	}
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perimeter())
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
