package main

import "fmt"

// used when there are multiple conditions
func main() {

	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// the program checks conditions in order until it finds a true case
	// the first true condition executes, and remaining conditions are ignored.
}
