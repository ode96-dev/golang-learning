package main

import "fmt"

/*
Pointers
- variables that store memory address of other variables.
- instead of passing a value directly to a function, we can pass their memory address
- can be used to allocate memory dynamically.
- helps:
  - in managing memory,
  - avoiding copies of large objects, and manipulating values directly

- Declared by placing an asterisk before the type, indicating that the variable is a pointer to the specified type.
- the default value of a pointer is nil, meaning it doesn't point to any memory location
- Go does not support pointer arithmetic directly

Syntax

	var p *int - p is a pointer to an integer
*/

// Passing by Pointer
func modifyValue(x *int) {
	*x = 100 // dereferencing the pointer to modify the value
}

// Pointers to Structs
type Person struct {
	name string
	age  int
}

func updateName(p *Person) {
	p.name = "Alice" // modifying struct field through a pointer
	p.age = 30
}

// Pointer to Array
func modifyArray(arr *[3]int) {
	arr[0] = 10 // modifying array via pointer
}

func main() {
	/*
		Declaring a Pointer
		- var p *int - declares a pointer to an integer
		- &a - & (address operator) used to get the memory address of variable a
		- *p - * (dereference operator) used to get the value stored at memory address the pointer points to.
	*/
	var a int = 58
	var p *int = &a // p is a pointer to the address of a

	fmt.Println("Address of a:", &a)       // memory address of a
	fmt.Println("Pointer p:", p)           // memory address stored in p
	fmt.Println("Value at pointer p:", *p) // value at the address

	/*
		Using Pointers with Functions
		- using pointers to modify a variable inside a function is a primary
		- modifyValue(&x) passess the address of x to the function
		- *x is used to dereference the pointer and modify the value of x
	*/
	x := 30
	fmt.Println("before:", x)
	fmt.Println("before:", &x)
	modifyValue(&x) // passing the address of x to the function
	fmt.Println("after:", x)
	fmt.Println("after:", &x)

	/*
		Pointers to Structs
		- we can use pointers to manipulate structs
		- passing a pointer to struct allows the function to modify the struct directly
	*/
	y := Person{name: "John", age: 25}
	fmt.Println("before:", y)

	updateName(&y) // passing to struct

	fmt.Println("after:", y)

	/*
		Pointers to Arrays and Slices
		- since slices are already reference types (they internally contain a pointer to the underlying array)
		- often, we do not need explicit pointers to slices
	*/
	arr := [3]int{1, 2, 3}
	fmt.Println("before:", arr)
	modifyArray(&arr) // passsing th address of the array
	fmt.Println("after:", arr)

	/*
		The new and make Functions in Go
		- new(type) - allocates memory for a variable of the specified type and returns a pointer to it
		- make(type, size) - allocates and initializes slices, maps, or channels, but does not return a pointer
	*/
	z := new(int)
	fmt.Println(*z)
	*z = 42
	fmt.Println(*z)

	// Using make
	m := make(map[string]int)
	m["age"] = 50
	fmt.Println(m)

	/*
		Nil Pointers and Pointer Safety
		- a pointer is nil by default if not initialized
		- it is important to check for nil pointers to avoid deferencing them, causing runtime errors
	*/
	var c *int // c is nil by default
	if c != nil {
		fmt.Println(*c)
	} else {
		fmt.Println("c is nil")
	}

}
