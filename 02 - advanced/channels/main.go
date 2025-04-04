package main

import (
	"fmt"
	"time"
)

/*
Channels in Go
- a channel is a built-in mechanism in Go that allows goroutines to communicate with each other
- it is a typed conduit through which you can send and receive values between concurrent goroutines
- the are:
	- safe, avoids race condition thus no explicit locking
	- efficient, used for concurrent communication
	- synchronized, blocks until the sender/receiver is ready
*/

// Channel Direction
// send only channel
func sendData(ch chan<- int) {
	ch <- 100
}

// receive only channel
func receiveData(ch <-chan int) {
	fmt.Println(<-ch)
}

func main() {
	/*
	   Declaring and Using Channels
	*/
	// declaring a channel
	var ch chan int
	fmt.Println(ch)

	// creating a channel with make() - by default a channel is nil and must be initialized by make()
	ch2 := make(chan int) // creates an int channel
	fmt.Println(ch2)

	go func() {
		ch2 <- 120 // send data to the channel
	}()

	number := <-ch2 // receive data from the channel
	fmt.Println(number)

	/*
	  Unbuffered vs Buffered Channels
	  - if buffer is full, sending blocks until space is available
	  - buffer channels store values until read.
	  - if buffer is empty, receiving blocks until data is available
	*/
	// Unbuffered Channels - synchronous communication - requires sender and receiver to ready at the same time
	ch3 := make(chan int) // unbuffered channel

	go func() {
		ch3 <- 10 // send (blocks until receiver is ready)
	}()

	fmt.Println(<-ch3) // receive (blocks until receiver sends)

	// Buffered Channels - asynchronous communication - allows sending without an immediate receiver, up to a certain capacity
	ch4 := make(chan string, 2) // buffered channel with capacity 2

	ch4 <- "Alice"
	ch4 <- "John"

	fmt.Println(<-ch4) // receive first name
	fmt.Println(<-ch4) // receive second name

	/*
		Closing a Channel
		-  a closed channel means no more value will be sent
		- closing a channel does not delete its values
		- reading after the channel is empty return zero values
		- sending to a closed channel causes a panic
	*/
	// Checking if a Channel is Closed
	ch5 := make(chan int, 2)

	ch5 <- 1
	ch5 <- 2

	close(ch5)

	for {
		v, ok := <-ch5

		if !ok {
			fmt.Println("channel closed")
			break
		}
		fmt.Println(v)
	}

	/*
		Iteration with range
	*/
	ch6 := make(chan string, 3)

	ch6 <- "Peter"
	ch6 <- "Alice"
	ch6 <- "John"

	close(ch6) // must close channel before using range

	for num := range ch6 { // range stops when channel is closed - this avoids the need for a manual check with ok
		fmt.Println(num)
	}

	/*
		Channel Direction
		- we can restrict channels to only sending and receiving
	*/
	ch7 := make(chan int)
	go sendData(ch7)
	receiveData(ch7)

	/*
		Select Statement - multiplexing channgel
		- select is used to wait on multiple channels
		- first channel executes first
		- no blocking since select is non-deterministic
	*/
	ch8 := make(chan string)
	ch9 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch8 <- "hello from ch8"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch9 <- "hello from ch9"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch8:
			fmt.Println(msg1)
		case msg2 := <-ch9:
			fmt.Println(msg2)
		}
	}
}
