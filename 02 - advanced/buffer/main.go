package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

/*
Buffers in Go
- buffer is a temporary storage area used to hold data before processing.
- allows efficient handling of data streams by storing and managing intermediate data before it is needed.
uses in Go:
	- buffered channels - storing values before retrieval
	- bytes.Buffer - efficient manipulation of byte slices
	- bufio package - buffered I/O operations
*/

func main() {
	/*
		Buffered channel usage
		- values are stored temporarily until retrieved
		- sender does not block until buffer is full
		- receiver does not block until buffer is empty
	*/
	ch := make(chan string, 2) // create buffered channel with capacity 2

	ch <- "message 1"
	ch <- "message 2"

	fmt.Println(<-ch) // read first message
	fmt.Println(<-ch) // read second message

	/*
		Blocking behavior in buffered channels
		- buffered channels block sending when full and block receiving when empty
	*/
	ch1 := make(chan int, 2) // capacity 2

	ch1 <- 1
	ch1 <- 2
	// ch1 <- 3 - this would be blocked because the buffer is full

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	/*
		bytes.Buffer for efficient string manipulation
		- provides an efficient way to work with strings and byte slices
		- using bytes.Buffer instead of string concatenation
			- efficient, avoids unnecessary memory allocations
			- fast, uses an internal buffer instead of copying strings
			- flexible, supports reading, writing, and manipulating byte slices
		- efficient string concatenation compared to str1+str2
		- utilizes internal memory efficiently instead of creating multiple copies
	*/
	var buffer bytes.Buffer // initialized a buffer

	buffer.WriteString("Hello, ")
	buffer.WriteString("Go! ")

	fmt.Println(buffer.String()) // convert buffer to string

	/*
		bufferio package package for Buffered I/O
		- bufio package provides buffered readers and writers for optimized I/O operations
		* bufio optimizes file reading/writing - reducing system calls
		* buffered input - read chunks instead of line by line
		* improved performance - especially for large data streams
	*/
	reader := bufio.NewReader(strings.NewReader("hello, Go World!\n Welcome to buffers."))

	line, _ := reader.ReadString('\n') // read until new line, \n
	fmt.Println(line)                  // output: Hello, Go

	/*
		Buffered Writer
		- buffered writing reduces file system calls
		- flush ensures data is written immediately
	*/
	file, _ := os.Create("output.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("buffered writing in Go\n")
	writer.Flush() // ensure all data is written to file

	fmt.Println("data is written to file...")

}
