package main

import "fmt"

func main() {
	number := 10

	if number%2 == 0 {
		fmt.Println("even number")
	} else {
		fmt.Println("odd number")
	}

	// if the condition number%2 == 0 is true, "even number" is printed
	// otherwise, the else block is executed
}
