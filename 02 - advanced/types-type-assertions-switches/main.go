package main

import "fmt"

/*
Type, Type Assertion, and Type Switches
- Go is strongly typed but also provides dynamic handling of typing through interfaces, type assertions, and switches
- Types must be defined or inferred
*/

// Using Type Switch
func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("integer:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}

}

func main() {

	/*
		Type Assertion
		- used to extract underlying concrete value from an interface{} type
		- allows you to convert an interface to a specific type
		Syntax
		- value, ok :=interfaceValue.(ConcreteType)
			- if ok == true, assertion succeeds.
			- if ok == false, assertion fails, panic if ok not checked
	*/
	// Type assertion with interface{}
	var i interface{} = "hello"

	str, ok := i.(string) // type assertion

	if ok {
		fmt.Println("String value:", str)
	} else {
		fmt.Println("type assertion failed")
	}

	/*
		Type Switches
		- a type switch allows handling of multiple possible types dynamically
		- uses interfaces{} and switch to check underlying types
		Syntax
		switch v := x.(type){
		case int:
			handle int
		case string:
			handle string
		default:
			handle unknown type
		}
	*/
	// Using Type Switch
	checkType(100)
	checkType("GoLang")
	checkType(false)
	checkType(3.142)

}
