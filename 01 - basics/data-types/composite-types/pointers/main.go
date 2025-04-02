package main

import "fmt"

/*
- Pointers are powerful feature that allows you to work with memory addresses directly
- used to store memory addresses of a variable
- useful when one needs to pass a large amount of data to a function or when you need to modify the value of a variable inside a function
- operators: & (Pointer to), *(Dereferencing pointer)
*/

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}
func main() {

	x := 8700
	fmt.Println(x)  // the value of x
	fmt.Println(&x) // to tell us where 7 is stored

	y := &x // equal to the pointer of x
	fmt.Println(x, y)

	*y = 8000
	fmt.Println(x, y)

	toChange := "hello"
	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

}
