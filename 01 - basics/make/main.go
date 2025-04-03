package main

import "fmt"

/*
- make() function in Go is used to create slices, maps, channels dynamically.
- unlike new, it allocate and initializes these types but does not return a pointer
- allocates zero memory for the object
- returns an initialized object
Syntax
 1. for slices - make(type, size, capacity)
 2. for maps - make(map[keyType]valueType)
 3. for channels - make(chan type, bufferSize)
*/
func main() {
	/*
		- make() with Slices:
		- elements are initialized to zero
		- capacity allows dynamic expansion without reallocation
		- make() initializes a slice with a specified length and capacity
		- length - number of elements in use
		- capacity - total space available before reallocation
	*/
	nums := make([]int, 3, 5) // slice length of 3, capacity 5
	fmt.Println("Slice:", nums)
	fmt.Println("Length:", len(nums))
	fmt.Println("Capacity:", cap(nums))

	/*
		- make() with maps:
		- make() is required to initialize map before using it
		- make(map[string]int) initializes an empty map
		- no need to specify size, Go handles growth dynamically
	*/
	scores := make(map[string]int)
	scores["Alice"] = 4
	scores["Bob"] = 0
	fmt.Println("Map", scores)

	/*
		TODO:CHANNELS
		- make() with channels:
		- make() is required to initialize channels
		- channels can be buffered (with a size) or unbuffered
		- unbuffered channels (make(chan int)) block until received
		- buffered channels (make(chan int, size)) store messages temporarily
	*/
	ch := make(chan int, 2)
	ch <- 10
	ch <- 30
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
