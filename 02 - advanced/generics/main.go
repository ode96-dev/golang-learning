package main

import (
	"fmt"
)

//TODO
/*
Generics
- provide ability to write functions, types, and data structures
- code reusability - write one implementation for different types
- type safety - compile-time checking ensures no invalid operations on types
- efficiency - avoiding code duplication and maintain optimized performance

Syntax
- type parameters in functions
  - we define generics in Go by introducing type parameters using square brackets [] after the function of type name
*/
func print[T any](value T) { // T is a type parameter which is replaced by a concrete type when the function is called
	fmt.Println(value)
}

// type constraint with comparable - used to ensure a function works only with types that can be compared using == or :=, use the comparable constraint
func findMax[T comparable](a, b T) T {
	if a == b {
		return a
	}

	return b
}

// defining a generic struct
type Pair[T any] struct {
	First  T
	Second T
}

// defining a generic map type
type Map[K comparable, V any] map[K]V

// Generic Constraint with Interface
type Stringer interface {
	String() string
}

// define a generic function that accepts only types that implement the Stringer interface
func printString[T Stringer](value T) {
	fmt.Println(value.String())
}

type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

// multiple type parameters
func swap[T1, T2 any](a T1, b T2) (T2, T1) {
	return b, a
}

func main() {
	// basic syntax, and type inferrence
	print(5)       // works with int
	print("hello") // works with string
	print(false)   // works with bool

	/*
		Type constraints in Go
		- restrict what types can be used for a given type parameter
			- any v specific constraints
				- any means no restriction on the type
				- you can define more specific constraints for more control over allowed types
	*/
	// type constraint with comparable - used to ensure a function works only with types that can be compared using == or :=, use the comparable constraint
	fmt.Println(findMax(4, 10))
	fmt.Println(findMax("apple", "oranges"))

	/*
		Generic data structures in Go
		- you can define generic types (structs, maps, slices) to work with any type
	*/
	// generic struct example - Pair[T any], defines a generic struct that can hold two values of any type
	pair1 := Pair[int]{First: 1, Second: 2}
	pair2 := Pair[string]{First: "Go", Second: "Lang"}

	fmt.Println(pair1)
	fmt.Println(pair2)

	// generic map - Map[K comparable, V any]defines a generic map where K is the key type (restricted to comparable types) and V is the value type
	stringMap := Map[string, int]{"Alice": 18, "John": 10}
	fmt.Println(stringMap)

	intMap := Map[int, string]{1: "one", 2: "two"}
	fmt.Println(intMap)

	/*
		Constraints with interfaces in Go
		- we can combine interfaces and generics to create more complex constraints for your functions and types
		- printString() can accept only types that implements the Stringer interface which has a String() method
	*/
	p := Person{Name: "John"}
	printString(p) // works because person implements Stringer

	/*
		Generic functions with multiple type parameters
		- we can define functions with multiple type parameters
	*/
	// multiple type parameters
	a, b := swap(1, "Golang")
	fmt.Println(a, b) // 1 hello

	x, y := swap("Programming", 3.142)
	fmt.Println(x, y)

}
