package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Network Request Handling
func mockAPIRequest(name string, ch chan string) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond) // simulate delay
	ch <- fmt.Sprintf("%s response received", name)
}

func main() {
	/*
		use case - API Requests
		- ensures the first API to respond gets handled immediately
		- prevents excessive waiting using time.After
	*/
	ch1 := make(chan string)
	ch2 := make(chan string)

	go mockAPIRequest("API-1", ch1)
	go mockAPIRequest("API-2", ch2)

	select {
	case res := <-ch1:
		fmt.Println(res)
	case res := <-ch2:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout: no API response within 2s")
	}
}
