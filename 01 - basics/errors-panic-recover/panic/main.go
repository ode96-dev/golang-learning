package main

import "fmt"

/*
panic() in Go
- used when program encounters an unrecoverable error
- it stops normal execution and starts unwinding the stack, calling deferred functions
- use only for fatal errors like corrupt memory or missing critical files
- should not be used for normal error handling
*/

// using panic()
func riskyOperation() {
	fmt.Println("Starting operation...")
	panic("something went terribly wrong!!") //terminate the program
}

func main() {
	fmt.Println("before panic")
	riskyOperation()
	fmt.Println("after panic") // this will never execute because execution stops immediately after panic()
}
