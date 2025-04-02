package main

import "fmt"

/**
DATA TYPES
- each variable has a type defined, holding values only for that type
- there two types, composite type and basic types.
- Allowed key types for map is any that the equality operator is defined: booleans, numbers, strings, arrays, pointers, structs, interafaces (as long as the dynamic type supports equality)
	- Structs, Maps, Functions are invalid key types because == is not defined for them
- Allowed value types: ANY

**/

func main() {
	fmt.Println("======BASIC TYPES======")
	var num int = 10                      // Integer
	var pi float64 = 3.14                 // Floating point number
	var isGoFun = true                    // Boolean
	var message string = "Go is awesome!" // String
	var letter rune = 'A'                 //rune (unicode character)

	fmt.Println(num, pi, isGoFun, message, letter)

}
