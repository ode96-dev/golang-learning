package main

import "fmt"

/*
*
- A variable is the name given to a memory location to store a value of a specific type.
- "var" keyword is used to declare a variable with a specific type, can be modified
- := allows short-hand variable declaration
- "const" keyword defines a constant which cannot be modified
- the syntax for declaring variables is:
	1. using var keyword - var variableName type = value
	2. using := sign - variableName := value
	3. using cont keyword - const CONSTNAME type = value
- Difference between := and var:
	var keyword
		- can be used inside and outside of functions
		- allows for declaration and value assignment done separately
	:=
		- can only be used inside functions
		- variable declaration and value assingment must be done in the same line.
- Naming Rules:
	1. must start with a letter or an underscore character
	2. cannot start with a digit
	3. alphanumeric characters and underscores
	4. case sensitive
	5. no limit to the length of a variable name
	6. cannot contain spaces
	7. cannot be any of Go keywords

	Techniques to making your code readable:
	1. Camel Case
	myName = "Peter"
	2. Pascal Case
	MyName = "Peter"
	3. Snake Case
	my_name = "Peter"
- constants should be uppercased to differentiate from variables
- constants can be declared both in and outside of a function

*/

func main() {

	var name string = "Alice" // explicit variable declaration
	age := 29                 //short-hand declaration of variables in Go. here, Go will infer the type.

	// typed constants - declared with a type
	const GENDER string = "Female" //constant, cannot be changed

	// when declaring a variable without an initial value, Go sets the default value to the assigned type
	var yourName string
	var yourAge int
	var isFemale bool

	//Go allows to assign values to a variable after their declaration. this helps in cases where values is unknown
	var profession string
	profession = "Engineer"

	//untyped constants - declared without a type
	const NAME = "Peter"

	// multiple constants declaration - can be grouped together into a block for readability
	const (
		X string = "Hello!"
		Y        = 3.42
		Z bool   = true
	)

	fmt.Println("====with initial value===")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(GENDER)

	fmt.Println("====without initial value===")
	fmt.Println(yourName)
	fmt.Println(yourAge)
	fmt.Println(isFemale)

	fmt.Println("====assigning value after declaration===")
	fmt.Println(profession)

	fmt.Println("====typed constant===")
	fmt.Println(GENDER)

	fmt.Println("====untyped constant===")
	fmt.Println(NAME)

	fmt.Println("====multiple constant declaration===")
	fmt.Println(X, Y, Z)

}
