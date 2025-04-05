package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Mutex in Go
- a mutex (mutual exclusion) is a synchronization mechanism that prevents multiple goroutines from accessing a shared resource at the same time
- prevents race condition (data corruption from concurrent access)
- ensures only one goroutine modifies a shared resource at a time
- simple and efficient for protecting critical sections
Syntax
	var mu sync.Mutex - declaring a mutex
*/

/*
Using Mutex to Prevent Data Races
*/
var counter int
var mu sync.Mutex // declare a mutex

func increment() {
	for i := 0; i < 5; i++ {
		mu.Lock()   // lock before modifying the shared resource
		counter++   // critical section
		mu.Unlock() // unlock after modification
		time.Sleep(time.Millisecond * 10)
	}
}

// Using sync.RWMutex for Read-Write Locks
var data int
var rw sync.RWMutex

func readData(id int) {
	rw.RLock() // read lock (multiple readers allowed)
	fmt.Printf("Goroutine %d reading: %d\n", id, data)
	time.Sleep(time.Millisecond * 500)
	rw.RUnlock() // release read lock
}

func writeData(value int) {
	rw.RLock() // exclusive write lock - blocks readers and writers
	fmt.Println("writing data:", value)
	data = value
	time.Sleep(time.Second)
	rw.RUnlock() // release write lock
}

// use case - Safe Bank Account Transactions
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (acc *BankAccount) Deposit(amount int) {
	acc.mu.Lock()
	acc.balance += amount
	fmt.Println("Deposited:", amount, "new balance:", acc.balance)
	acc.mu.Unlock()
}
func (acc *BankAccount) Withdraw(amount int) {
	acc.mu.Lock()
	if acc.balance >= amount {
		acc.balance -= amount
		fmt.Println("withdrew:", amount, "new balance:", acc.balance)
	} else {
		fmt.Println("insufficient funds")
	}
	acc.mu.Unlock()
}

func main() {

	/*
	   Using Mutex to Prevent Data Races
	   - deadlock is common issue where a goroutine locks and is not released, ensure to unlock after locking
	*/
	go increment()
	go increment()

	time.Sleep(time.Second)
	fmt.Println("final counter:", counter)

	/*
		defer to ensure unlocking
		- safer because the lock is always released at the end of the function
		func increment(){
			mu.Lock()
			defer mu.Unlock() - ensures unlock even if an error occurs
			counter++
		}
	*/

	/*
		Using sync.RWMutex for Read-Write Locks
		- allows multiple readers but only one writer
		- improves performance when many goroutiines only need to read data
		- allows concurrent reads, but writes are exclusive
			- multiple readers can access at the same time
			- writer blocks all others while writing
	*/
	go readData(1)
	go readData(2)
	go writeData(42)

	time.Sleep(3 * time.Second)

	/*
		use case - Safe Bank Account Transactions
		- prevents concurrent balance updates from corrupting data
	*/
	acc := BankAccount{balance: 100}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		acc.Deposit(50)
		wg.Done()
	}()

	go func() {
		acc.Withdraw(30)
		wg.Done()
	}()

	go func() {
		acc.Withdraw(500)
		wg.Done()
	}()

	wg.Wait() // wait for all goroutines to finish
}
