package main

import (
	"fmt"
	"math"
)

/*
Interfaces
- an interface in Go defines a set of method signatures but does not implement them.
- a type implements an interface implicitly by defining all the methods declared in the interface
- interfaces provides abstration - allows different types to be used interchangebly
- interfaces are satisfied implicitly in Go - a struct automatically implements an interface if it has all the required methods
Syntax
- type InterfaceName interface {
	Method1(param1, Type) ReturnType
	Method 2(param2, Type) ReturnType
}
*/

// Implementing an Interface
type Animal interface {
	Speak() string
}

// Define a string for a Dog
type Dog struct {
	Name string
}

// Implement the Speak() method for Dog
func (d Dog) Speak() string {
	return "woof!"
}

// multiple implementation of an interface - polymorphism, different types implementing the same interface
// Different Animals Implementing Speak()
type Cat struct{}
type Cow struct{}

func (c Cat) Speak() string {
	return "meow"
}
func (c Cow) Speak() string {
	return "mooo"
}

// Interfaces with Multiple Methods - an interface can have multiple methods and a type must implement all of them
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width * r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Empty Interface (interface{})
func printValue(val interface{}) {
	fmt.Println("Value:", val)

}

func main() {
	// Basic interface implementation
	var myDog Animal = Dog{Name: "Buddy"} // implicit implementation
	fmt.Println(myDog.Speak())            // calls the Speak() method

	// multiple implementation of an interface
	animals := []Animal{Dog{Name: "Buddy"}, Cat{}, Cow{}} // slice of different Animal implementation

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	// Interfaces with Multiple Methods
	var shapes = []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
	}

	for _, shape := range shapes {
		fmt.Printf("Area: %2f, Perimeter: %2f\n", shape.Area(), shape.Perimeter())
	}

	// Empty Interface (interface{}) - an empty interface can hold any type, making it useful in handling generic types
	printValue(40)
	printValue("hello")
	printValue(false)

	// Type Assertion with Interfaces - allows extracting the concrete type from an interface{}
	var x interface{} = "Alice"

	str, ok := x.(string)

	if ok {
		fmt.Println("String:", str)
	} else {
		fmt.Println("failed type assertion")
	}

}
