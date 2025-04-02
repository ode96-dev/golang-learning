package main

import "fmt"

/*
   - Go has only one loop
   - has three components:
	1. the init statement - executed before the first iteration
	2. the condition expression - evaluated before every iteration
	3. the post statement - executed at the end of every iteration
  - range keyword in Go is used to iterate over different data structures
*/

func main() {
	// classic syntax
	for i := 1; i <= 5; i++ {
		fmt.Println("iterations", i)
	}

	/*
	   - i := 1 initializes i to 1
	   - i <= 5 runs the loop while i is ≤ 5
	   -  i++ increments i after each iteration
	*/

	//for loop without initialization & post statement
	// useful when the loop condition is dependent on external factors

	x := 1
	for x <= 5 {
		fmt.Println("iteration:", x)
		x++
	}

	// infinit loop, for... without condition
	// for {
	// 	fmt.Println("this is an infinite loop")
	// }

	/*
	   - using break to exit the loop
	   - loop breaks i == 5
	*/
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("loop terminated at: ", i)
			break
		}
		fmt.Println("iteration: ", i)
	}

	/*
	   - using continue to skip iteration
	   - skips iteration when i == 3
	*/
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("skipping iteration", i)
			continue
		}
		fmt.Println("iteration", i)
	}

	/*
	   - nested for loop
	   - useful for 2D data processing, grid-based application, matrix operations
	*/
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

	/*
	   - looping over a slice/array
	   - fruit is the value at index
	*/
	fruits := []string{"Apple", "Banana", "Orange"}

	for i, fruit := range fruits {
		fmt.Println("Index: ", i+1, "Fruit: ", fruit)
	}

	/*
	   - looping over a map(key-value pair)
	   - useful in processing json data
	*/
	user := map[string]string{
		"name":   "Alice",
		"role":   "Admin",
		"status": "Active",
	}

	for key, value := range user {
		fmt.Println(key, ":", value)
	}

	/*
	   - looping over a string
	*/
	text := "GoLang"

	for i, char := range text {
		fmt.Printf("Index: %d, Character: %c\n", i+1, char)
	}

	/*
	   - sample application
	   - brute-force password cracker
	*/
	password := "1234"

	for attempt := 0; attempt <= 9999; attempt++ {
		if fmt.Sprintf("%04d", attempt) == password {
			fmt.Println("password cracked: ", attempt)
			break
		}
	}
}
