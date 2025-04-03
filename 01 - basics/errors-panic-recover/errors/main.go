package main

import (
	"errors"
	"fmt"
)

/*
Errors in Go (error Interface)
- errors in Go are values that implement the error interface
- a function that will fail typically returns an error as its last return value
- nil means no error recorded
*/

// function that may return an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	// Returning an error
	result, err := divide(10.10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result", result)

}
