package main

import "fmt"

/*
*ARRAY
- A collection of of elements of the same type with a fixed size defined when the array is created.
- Declared by either := or var keyword
- if an array or any of its element has not been initialized, then it is assigned the default value of its type
*/
func main() {
	fmt.Println("======Arrays======")
	var numbers = [3]int{10, 20, 30}               //fixed sized array
	var names = [3]string{"Peter", "Paul", "Mary"} //fixed sized array
	options := [5]bool{true, false, true, false, false}
	fmt.Println(numbers, names, options)

	// accessing elements of an array
	fmt.Println(names[1])

	// changing elements of an array
	names[0] = "John"
	fmt.Println(names)

	// finding length of an array - using len() function
	arr2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr2))

}
