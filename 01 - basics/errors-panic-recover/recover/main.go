package main

import "fmt"

/*
recover() in Go
- used inside of defer function to catch panic and prevent program crash
- retrieves the error message and allows the program to continue running
- program continues execution after recover()
- recover() should only be used insde defer function
*/

// Using recover()
func safeOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Performing risky operation...")
	panic("Critical failure!") // Triggers panic
	fmt.Println("This line will not execute")
}

func main() {
	fmt.Println("Before safeOperation")
	safeOperation()
	fmt.Println("After safeOperation") // Execution continues after recovery
}
