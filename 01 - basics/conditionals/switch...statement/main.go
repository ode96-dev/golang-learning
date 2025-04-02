package main

import "fmt"

/*
- the expression is evaluated once
- the value of the switch expression is compared with the values of each case
- if there is a match, the associated block of code is executed
- the default keyword is optional. it specifies some code to run if there is no case match
*/

func main() {
	// switch without condition
	// switch can be used without a variable

	num := -5

	switch {
	case num < 0:
		fmt.Println("negative number")
	case num == 0:
		fmt.Println("zero")
	default:
		fmt.Println("positive number")
	}

	// switch with multiple case values
	// useful for grouping multiple cases together

	grade := "B"

	switch grade {
	case "A", "B": // shares the same case block
		fmt.Println("excellent performance")
	case "C", "D":
		fmt.Println("needs improvement")
	default:
		fmt.Println("invalid grade")
	}

	// switch with Failthrough
	// by default, Go does not fall through cases
	// you need to explicitly use fallthrough if you want the next case to run
	// fallthrough forces execution of the next case even if the condition is false
	// without fallthrough, only the first match case runs

	num2 := 2

	switch num2 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three, fall through executed")
	default:
		fmt.Println("unknown number")
	}

}
