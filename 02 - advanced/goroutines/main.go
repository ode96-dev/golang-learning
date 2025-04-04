package main

import (
	"fmt"
	"sync"
	"time"
)

/*
TODO: how to resolve race condition in Goroutines
Goroutine
- is a lightweight thread of execution in Go.
- allows concurrent execution of functions without blocking the main thread
  - it is efficient, uses much less memory than OS
  - it is fast, you can spawn thousands of goroutines efficiently
  - automatic management - Go's runtime schedules them efficiently
*/

// Creating a Goroutine
func printMessage() {
	fmt.Println("Good Go Goroutines")
}

// Goroutines with Arguments
func greet(name string) {
	fmt.Println("Hello", name)
}

// using sync.WaitGroup
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // marks the goroutine as done
	fmt.Println("worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Println("worker %d finished\n", id)
}

// using channels
func worker2(ch chan string) {
	ch <- "Work done" // send data to the channel
}

func main() {
	/*
		Creating a Goroutine
		- created by prefixing a function call with go
		- go printMessage() runs printMessage() concurrently
		- time.Sleep(time.Second) prevents the program from exiting before the goroutine executes
	*/
	go printMessage() // run function in a separate goroutine

	time.Sleep(time.Second) // wait for Goroutine to finish

	fmt.Println("Main function finished")

	/*
		Goroutines with Anonymous Functions

	*/
	go func() {
		fmt.Println("Go with anonymous function")
	}()

	time.Sleep(time.Second) // wait for the Goroutine

	/*
		Goroutines with Arguments
	*/
	go greet("Alice")       //passing argument to Goroutine
	time.Sleep(time.Second) // wait for Goroutine

	/*
		Common Issue in Goroutines - Race Condition
		- happens when multiple goroutines access shared data concurrently
	*/
	for i := 1; i <= 5; i++ {
		go fmt.Println("Count:", i) // results to race condition
	}
	time.Sleep(time.Second)

	/*
		Synchronizing Goroutines
		- Goroutines do not wait for each other
		- we need sync mechanisms like WaitGroups
	*/
	// using sync.WaitGroup
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed")

	/*
		Communication Between Goroutines
		- Goroutines communicate using channels
	*/
	ch := make(chan string)
	go worker2(ch) // start worker Goroutine
	msg := <-ch
	fmt.Println(msg)
}
