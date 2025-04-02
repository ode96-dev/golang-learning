package main

import "fmt"

/*
*
- a struct (structure) is utilized to create a collection of members of different data types, into a single variable
- unlike slices and arrays, struct store multiple values of different types into a single variable.
- useful for grouping data together to create records
*
*/

// declaring a struct
type Person struct {
	name     string
	age      int
	isFemale bool
}

func main() {
	person1 := Person{name: "Alice", age: 29, isFemale: true}

	// accessing struct members
	fmt.Println(person1.name, "is", person1.age, "years old.", "She is", person1.isFemale)
}
