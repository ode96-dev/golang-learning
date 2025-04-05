package main

import (
	"fmt"
	"time"
)

/*
select in Go
- is a powerful multiplexing tool that allows a goroutine to wait on multiple channel operations simultaneously.
- enables non-blocking communication between goroutines, making it a key feature in Go's concurrency model.
	- listens to multiple channels at once
	- executes the first ready case (random selection if multiple are ready)
	- prevents blocking if used with a default case
	- efficient for handling concurrency/timeouts
Syntax
	select {
	case val := <-ch1: --received value from ch1
	case ch2<-10: --sent value to ch2
	default: --default case if no channels are ready
	}
*/

func main() {
	/*
		Handling multiple channels with select
	*/
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "data from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "data from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}

	/*
		using select timeouts
		- prevents indefinite blocking
		- gracefully handles long-running operations
	*/
	ch3 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch3 <- "response from goroutine"
	}()

	select {
	case msg := <-ch3:
		fmt.Println("received:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout! no response received")
	}

	/*
		select with default with non-blocking operations
		- adding default ensures that select does not block if no channels are ready
		- avoids blocking if no channel is ready
		- useful for checking channels without waiting
	*/
	ch4 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch4 <- "Alice"
	}()

	select {
	case val := <-ch4:
		fmt.Println("received:", val)
	default:
		fmt.Println("no data available, moving on...")
	}

	/*
		using select in an infinite loop
		- to continously listen to multiple channels, we use select inside a for loop
		- keeps listening until 5 second inactivity timeout
	*/
	ch5 := make(chan string)
	ch6 := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch5 <- "ping from ch5"
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch6 <- "ping from ch6"
		}
	}()

	for {
		select {
		case msg1 := <-ch5:
			fmt.Println(msg1)
		case msg2 := <-ch6:
			fmt.Println(msg2)
		case <-time.After(5 * time.Second):
			fmt.Println("no activity for 5 seconds. exiting")
			return
		}
	}

}
