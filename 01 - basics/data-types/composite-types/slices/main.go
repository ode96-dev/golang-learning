package main

import "fmt"

/*
- similar to arrays, but more powerful and flexible
- used to store multiple values of the same type in a single variable
- unlike arrays, the length of a slice can grow and shrink whenever.
- different ways to create slices in Go:
	1. using []datatype{value} format
	2. using an array
	3. using make() function
- two functions can be used to return the length and capacity of a slice:
	1. len() - the number of elements in v
	2. cap() - the maximum length the slice can reach when resliced
*/

func main() {
	fmt.Println("======Slices, commonly used than arrays in Go======")
	slice := []int{1, 2, 3, 4, 5} //dynamic size
	slice = append(slice, 6)      // append new elements
	fmt.Println(slice)

	// using []datatype{value} format: syntax - slice_name := []datatype{value}
	sampleSlice := []int{1, 2, 3}
	fmt.Println(sampleSlice)

	//checking length and capacity
	sampleSlice2 := []string{"Go", "Slices", "Are", "Cool"}
	fmt.Println(sampleSlice2)
	fmt.Println(len(sampleSlice2))
	fmt.Println(cap(sampleSlice2))

	// using an array
	sampleArrayForSlice := [4]string{"Peter", "Paul", "Mary", "Jane"}
	sampleSlice3 := sampleArrayForSlice[0:2]
	fmt.Println(sampleSlice3)
	fmt.Println(len(sampleSlice3))
	fmt.Println(cap(sampleSlice3))

	// using make() function - slice_name := make([]type, length, capacity)
	// if capacity is not defined, it will be equal to length
	sampleSlice4 := make([]string, 4, 8)
	sampleSlice5 := make([]string, 4) // without capacity defined
	fmt.Println(sampleSlice4)
	fmt.Println(len(sampleSlice4))
	fmt.Println(cap(sampleSlice4))
	fmt.Println(len(sampleSlice5))
	fmt.Println(cap(sampleSlice5))

	// accessing, changing, appending slices
	fmt.Println(sampleSlice2[0]) // accessing slices using index number
	fmt.Println(sampleSlice2[1]) // accessing slices using index number

	fmt.Println(sampleSlice2[3])
	sampleSlice2[3] = "Powerful" // changing slices
	fmt.Println(sampleSlice2[3]) // changing slices

	sampleSlice = append(sampleSlice, 4, 5, 6, 7, 8) //appending slices: slice_name = append(slice_name, element1, element2,...)
	fmt.Println(sampleSlice)

	sampleSlice6 := append(sampleSlice, slice...) // appending one slice to another: slice_name = append(slice1, slice2...) - ... is necessary when appending one slice to another
	fmt.Println(sampleSlice6)

}
