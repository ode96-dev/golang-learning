package main

import "fmt"

// curly braces {} are required even if there is only one statement.

func main() {
	age := 21

	if age >= 21 {
		fmt.Println("You can buy alcohol yet!")
	}

	// block inside {} only runs if age >=21
}
