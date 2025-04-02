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

// Function Accepting Multiple Arguments
func sum3(numbers ...int) int {
	total := 0

	for _, num := range numbers {
		total += num
	}

	return total
}

// Higher-Order Functions (Functions as Arguments)
func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
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

	// Function Accepting Multiple Arguments
	fmt.Println("Total:", sum3(1, 2, 3, 4, 5))
	/*
	   - Notes:
	     1. * ...type will accept any number of type arguments
	*/

	// Anonymous Functions (Lambdas)
	square2 := func(n int) (result int) {
		result = n * n
		return
	}
	fmt.Println("Square:", square2(5))
	/*
		   - Notes:
		     1. used for short-lived, inline operations
			 2. used in event driven programming
			 3. used for callbacks
	*/

	// Higher-Order Functions (Functions as Arguments)
	add3 := func(a, b int) int { return a + b }
	divide2 := func(a, b int) int { return a / b }
	fmt.Println("Result:", applyOperation(10, 5, add3))
	fmt.Println("Result2:", applyOperation(10, 5, divide2))
	/*
		   - Notes:
		     1. functions can accept other functions as argument or returns a function
			 2. useful in creating middlewares in apps.
			 3. encapsulation, readability, reusability
	*/

}
