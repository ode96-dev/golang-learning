package main

import (
	"fmt"
	"runtime"
	"time"
)

//TODO
/*
Scheduler in Go
- part of the Go runtime that manages the execution of goroutines on available CPU cores.
- Go does not use OS threads directly for every goroutine; instead, it maps goroutines to a smaller number of OS threads, making execution efficient and scalable
 - lightweight goroutines instead of heavyweight threads
 - goroutines execution in parallel on multiple OS threads
 - uses a work-stealing algorithm for load balancing
 - optimized for high-performance concurrent applications.

Understanding the Go Scheduler’s Execution Model
- Go uses G-P-M model (Goroutine (a lightweight thread managed by the Go runtime), Processor (a logical processor that runs goroutines), Machine (an OS thread used to execute Go code)) to schedule goroutines
How it works:
1. Go runtime creates Goroutines (G)
2. Each goroutine is assigned to a Processor (P)
3. Processor (P) run on OS threads (M)
4. If a goroutine blocks (e.g. waiting for I/O), the scheduler moves another goroutines onto the thread
*/

/*
How Go schedules goroutines
*/
func worker(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", id, i)
		time.Sleep(time.Millisecond * 500) // simulates work
	}
}

/*
Preemptive scheduling in Go
- Go uses cooperative scheduling with preemption to ensure that long-running goroutines do not block others
*/
func busyLoop() {
	for i := 0; i < 1000000000; i++ { // very long loop
		if i%1000000 == 0 {
			runtime.Gosched() // yields to allow other goroutines to run
		}
	}
}

// blocking with sleep
func worker2(id int) {
	fmt.Printf("worker %d started\n", id)
	time.Sleep(time.Second * 2) // simulating I/O delay
	fmt.Printf("worker %d finished\n", id)
}

/*
use case - task scheduling in Go
*/
func scheduledTask() {
	for {
		fmt.Println("executing scheduled task:", time.Now())
		time.Sleep(2 * time.Second) // runs every 2 seconds
	}
}

func main() {

	/*
		How Go schedules goroutines
		- the scheduler efficient distributes goroutines among CPU cores
		- execution order is non-deterministic due to concurrent scheduling
	*/
	fmt.Println("number of cpus:", runtime.NumCPU()) // shows available cpu cores
	fmt.Println("number of goroutines before:", runtime.NumGoroutine())

	for i := 1; i <= 3; i++ {
		go worker(i) // spawn multiple goroutines
	}

	fmt.Println("number of goroutines after:", runtime.NumGoroutine())
	time.Sleep(time.Second * 3) // wait to let goroutines finish

	/*
		GOMAXPROCS
		- by default, Go uses all CPU cores, but you can control this using runtime.GOMAXPROCS()
		- controls how many OS threads run goroutines in parallel
		- defaults to runtime.NumCPU()
	*/
	runtime.GOMAXPROCS(2) // limit to 2 CPU cores
	fmt.Println("GOMAXPROCS set to:", runtime.GOMAXPROCS(0))

	/*
		Preemptive scheduling in Go
		- Go uses cooperative scheduling with preemption to ensure that long-running goroutines do not block others
		- runtime.Goshed() yields control to other goroutines
		- preemptive scheduling ensures fair execution
	*/
	go busyLoop()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("other work happening...")
			time.Sleep(time.Millisecond * 500)
		}
	}()
	time.Sleep(time.Second * 3)

	/*
			Work stealing in Go
			- work-stealing algorithm helps balance goroutines across processors
			1. Each Processor(P) maintains a local queue of goroutines
			2. when a Processor(P) is idle, it steals work from another Processor(P)
			3. this keeps the CPU fully utilized
		- prevents processors from sitting idle while others are overloaded
	*/

	/*
		Go scheduler and block operations
		- if a goroutine blocks I/O, network, or system calls, Go parks it and assigns another goroutine to he OS thread
	*/
	// blocking with sleep
	go worker2(1)
	go worker2(2)

	time.Sleep(time.Second * 2) // allows goroutines to complete

	/*
		use case - task scheduling in Go
		- runs tasks at scheduled intervals
		- ideal for background jobs
	*/
	go scheduledTask()           // background scheduler
	time.Sleep(10 * time.Second) // keep program running

}
