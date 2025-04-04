package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

/* TODO
context
- the context package in Go provide a way to manage timeouts, deadlines, and cancellation signals across API calls, goroutines, and distributed systems.
- why context:
	- graceful cancellation - stop goroutines when they are no longer needed
	- timeout handling - prevents requests from hanging indefinitely
	- deadline management - ensures tasks finish within a specific time
	- pass metadata - shares request-specific values across API calls
*/
// Context Creation and Cancellation
// worker function that listens for cancellation
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "stopped")
			return
		default:
			fmt.Println("worker", id, "is working")
			time.Sleep(time.Second)
		}
	}
}

// Context with Timeout
func task(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task Completed")
	case <-ctx.Done():
		fmt.Println("Task Cancellation", ctx.Err())
	}

}

// Context with Deadline
func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("work stopped:", ctx.Err())
			return
		default:
			fmt.Println("working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Function that retrieves a value from context
func processRequest(ctx context.Context) {
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		fmt.Println("no user id found")
		return
	}
	fmt.Println("processing request for user id:", userID)
}

// Context in an HTTP Server
func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "request completed")
	case <-ctx.Done():
		http.Error(w, "Request timed out:", http.StatusRequestTimeout)
	}
}

func main() {
	// Context Creation and Cancellation
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(3 * time.Second)
	fmt.Println("Stopping workers")
	cancel() // cancel all workers

	time.Sleep(time.Second) // allow goroutines to finish

	// Context with Timeout
	ctx2, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go task(ctx2)

	time.Sleep(3 * time.Second)

	// Context with Deadline
	deadline := time.Now().Add(2 * time.Second) // set deadline

	ctx3, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go work(ctx3)

	time.Sleep(3 * time.Second)

	// Passing Values with Context - context can store values e.g. user id, request id
	ctx4 := context.WithValue(context.Background(), "userID", 12345)
	processRequest(ctx4)

	// Context in an HTTP Server - in web apps, helps to cancel requests, set timeouts, and manage request-scoped values
	http.HandleFunc("/", handler)
	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
