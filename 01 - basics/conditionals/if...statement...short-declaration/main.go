package main

import "fmt"

func main() {
	if num := 42; num > 0 {
		fmt.Println("positive number")
	}

	// num := 42 is declared inside the if condition
	// here, num is only available inside the if block
}
