package main

import "fmt"

/*
Type Inference
- allows the compiler to automatically determine the type of a variable based on the assigned value
- eliminates the need to explicitly specify types in many cases, making code cleaner and more concise
- works with variables, function returns...
- Go still enforces strong typing
*/

// Global variable with inferred type
var message = "Hello, Go!"

// Inferred Constant Types
const PI = 3.142
const ANSWER = 420

// Function returns an inferred type (int)
func add(a, b int) int {
	return a + b
}

// Type Inference in Struct Fields
type Person struct {
	name  string
	age   int
	score float64
}

// Dynamic Type Inference
func printValue(v interface{}) {
	fmt.Println("Value:", v)

}

func main() {
	/*
		Basic type inference using := Operator
		- := Operator automatically infers the type based on the assigned value
			- it can only be used inside functions (not for global variables)
	*/
	name := "GoLang" //inferred as string
	age := 25        // inferred as integer
	pi := 3.142      // inferred as float64
	isGoFun := true  // inferred as bool

	fmt.Println(name, age, pi, isGoFun)

	/*
		Type Inference with var Keyword
		- when using var, Go can still infer types if no explicit type is provided
		- unlike :=, var works inside and outside functions
	*/
	var x, y, z = 10, 20.5, true // inferred as int, float64, bool
	fmt.Println(x, y, z)

	/*
		Type Inference in Constants
			- constants in Go support type inference
			- if no type is explicitly given, the compiler assigns the most appropriate type
			- constants remain untyped until they are assigned to a variable
			- thus, you can easily convert them when needed
	*/
	fmt.Println(PI, ANSWER)

	var pi2 float64 = float64(PI)
	fmt.Println(pi2)

	/*
		Type Inference with Function Return Values
			- when a function returns a value, Go infers the type of the returned expression
	*/
	result := add(5, 10) // result is inferred as int based on add() return type
	fmt.Println(result)

	/*
		Type Inference in Structs
			- Go can infer types for struct fields if they are assigned a value
	*/
	p := Person{"Alice", 30, 95.3} // inferred type fields
	fmt.Println(p.name, p.age, p.score)

	/*
		Type Inference with Slices and Maps
			- Go can infer types when declaring slices and maps
	*/
	// Type Inference with Slices
	numbers := []int{1, 2, 3, 4, 5}             // inferred as []int
	names := []string{"Alice", "Peter", "Mary"} // inferred as []string

	fmt.Println(numbers, names)

	// Type Inference with Maps
	ages := map[string]int{
		"Alice": 20,
		"Peter": 35,
		"Mary":  18,
	} // Inferred as map[string]int
	fmt.Println(ages)

	/*
		Type Inference with Interfaces (interface{})
			- If a value is assigned to an interface{}, Go infers its dynamic type at runtime
			- interface{} dynamically adapt to the assigned value's type
				- useful for generic functions that handle multiple types
	*/
	// Dynamic Type Inference
	printValue(42)          // inferred as int
	printValue("Hello Go!") // inferred as string
	printValue(3.142)       // inferred as float64

}
