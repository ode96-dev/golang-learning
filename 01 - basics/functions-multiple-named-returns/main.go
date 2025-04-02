package main

import "fmt"

/*
- Functions are fundamental blocks in Go.
- Enables reusability, readability, modularity.
- In Go, a function is a block of reusable code that performs a given task.
*/

/*
- Basic Syntax
	func functionName(parameters) returnType {
		return result
	}
*/

// function definition, Function without parameters
func greet() {
	fmt.Println("Hello, Go!")
}

// function with parameters, one
func greet2(name string) {
	fmt.Println("hello", name)
}

// Function with Multiple Parameters
func add(a int, b int) int {
	return a + b
}

// Function with return values, one
func square(n int) int {
	result := n * n
	return result
}

// Function with return values, multiple
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Named Return Values
func add2(a, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	// Function without parameters
	greet() // function call
	/*
		- Notes:
		1. greet(), a simple function printing out a message.
		2. no parameters or return values given.
	*/

	// function with parameters, one
	greet2("Brian") // function call
	/*
		- Notes:
		1. name string is a parameter that accepts a string
	*/

	// Function with Multiple Parameters
	result := add(5, 3)
	fmt.Println("Sum for add: ", result)
	/*
		- Notes:
		1. return type of int, function must return an integer
	*/

	// Function with return values, one
	fmt.Println("Square:", square(4))
	/*
		- Notes:
		1. return sends the computed value back to the caller.
	*/

	// Function with return values, multiple
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)
	/*
		- Notes:
		1. multiple return values helps to avoid global variables
		2. multiple return values helps to improve function clarity
		3. multiple return values helps to reduce errors in complex logic
	*/

	// Named Return Values
	fmt.Println("Sum 2:", add2(4, 6))
	/*
		   - Notes:
		     1. the variable sum is automatically return
			 2. named return values improve readability when return values are self-explanatory
	*/
}
