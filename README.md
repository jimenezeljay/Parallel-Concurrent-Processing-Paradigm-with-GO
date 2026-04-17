# Parallel-Concurrent-Processing-Paradigm-with-GO
This is a project working with a new language I'm learning: GO

# Parallel & Concurrent Processing Paradigm with Go
### Project 3 - CSC 310 Programming Languages
Language: Go (Golang) | Topic: Parallel Number Crunching with Goroutines & Channels

---

## Project Overview

This project implements a parallel prime number finder in Go, demonstrating the difference between single-threaded and multi-threaded computation using Go's built-in concurrency primitives - goroutines and channels.

---

## How to Run

go run main.go

---

## Go Documentation Reference

### 1. Go Basics - Variables & Functions
Covers: variable declaration, short declaration (:=), static typing, function syntax

Tour of Go - Basics: https://go.dev/tour/basics/1
Variables: https://go.dev/tour/basics/8
Functions: https://go.dev/tour/basics/4

Key concepts:
- var name string = "value" - explicit variable declaration
- := - short declaration, Go infers the type automatically
- func name(param type) returnType {} - return type goes at the end
- Go is statically typed - types cannot change after declaration

---

### 2. Loops & Slices
Covers: Go's for loop (the only loop), slices, range, append

Tour of Go - For loop: https://go.dev/tour/flowcontrol/1
Tour of Go - Slices: https://go.dev/tour/moretypes/7

Key concepts:
- for is Go's only loop - it replaces while too
- []int{1, 2, 3} - slice (Go's dynamic list/array)
- range - iterates over a slice, returns index and value
- append(slice, value) - adds an element to a slice
- _ - blank identifier, used to ignore a value

---

### 3. Goroutines
Covers: launching goroutines, concurrent execution, the go keyword

Tour of Go - Goroutines: https://go.dev/tour/concurrency/1
Effective Go - Goroutines: https://go.dev/doc/effective_go#goroutines

Key concepts:
- go functionName() - launches a function as a goroutine (runs concurrently)
- Goroutines are lightweight threads managed by Go's runtime
- Go can run thousands of goroutines simultaneously
- main() exits and kills all goroutines if it finishes before them - use channels to sync

---

### 4. Channels
Covers: creating channels, sending/receiving values, blocking, close()

Tour of Go - Channels: https://go.dev/tour/concurrency/2
Tour of Go - Buffered Channels: https://go.dev/tour/concurrency/3
Tour of Go - Range and Close: https://go.dev/tour/concurrency/4

Key concepts:
- make(chan int) - creates a channel that carries int values
- ch <- value - sends a value INTO a channel
- value := <-ch - receives a value FROM a channel
- Channels block - receiving waits until something is sent (no need for time.Sleep)
- close(ch) - signals no more values will be sent
- for n := range ch {} - receives from channel until it is closed

---

### 5. Performance Timing
Covers: measuring execution time with the time package

Go time package: https://pkg.go.dev/time

Key concepts:
- start := time.Now() - records current time
- elapsed := time.Since(start) - calculates time elapsed since start
- Used to compare single-threaded vs. parallel performance

---

### 6. Math Package
Covers: math.Sqrt() used in the prime checking optimization

Go math package: https://pkg.go.dev/math

Key concepts:
- math.Sqrt(float64(n)) - square root, used to optimize prime checking
- Only need to check divisors up to √n to determine if n is prime

---

### 7. fmt Package
Covers: printing output to the terminal

Go fmt package: https://pkg.go.dev/fmt

Key concepts:
- fmt.Println() - prints with a newline
- fmt.Printf() - formatted print, uses %d (int), %s (string), %f (float)

---

### 8. Go Modules & Project Structure
Covers: setting up a Go project

Go modules: https://go.dev/doc/code

Key concepts:
- go mod init module-name - initializes a Go module
- go run main.go - runs the program directly
- go build - compiles the program into an executable

---

## Concurrency Concepts Used in This Project

Goroutines - Lightweight concurrent threads - go worker(jobs, results)
Channels - Communication between goroutines - make(chan int, size)
Worker Pool - Fixed number of goroutines sharing work - for w := 0; w < workerCount; w++
Task Distribution - Splitting work across workers - Sending numbers to jobs channel
Synchronization - Collecting all results before finishing - Receiving from results channel

---

## Performance Results

Single-Threaded | 78,498 primes (up to 1,000,000) | ~330ms
Parallel | 78,498 primes (up to 1,000,000) | TBD

---

## Additional Learning Resources

Official Go Tour (interactive): https://go.dev/tour
Go by Example: https://gobyexample.com
Effective Go (best practices): https://go.dev/doc/effective_go
Go Standard Library: https://pkg.go.dev/std
Go Playground (run Go in browser): https://go.dev/play

---

That's everything! The parallel time result is TBD - fill that in tomorrow once we finish Step 6. See you then! 💪
