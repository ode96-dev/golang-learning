package main

import "fmt"

func main() {
	age := 20
	hasID := true

	if age >= 18 { // first condition
		if hasID { // second condition inside the first if
			fmt.Println("you are allowed to enter")
		} else {
			fmt.Println("you need an ID to enter")
		}
	} else {
		fmt.Println("chill!!! you are underage")
	}

	// first, we check if age >= 18
	// if true, we enter the inner if block and checks if hasID == true
	// if both are true, you are allowed to enter
	// if hasID is false, you need an ID to enter
	// if age < 18, chill!!! you are underage

	// utilizing logical operators, avoids nesting for readability
	username := "admin"
	password := "secure123"
	isActive := true

	if username == "admin" && password == "secure123" && isActive {
		fmt.Println("Access is granted.")
	} else {
		fmt.Println("Access denied.")
	}
}
