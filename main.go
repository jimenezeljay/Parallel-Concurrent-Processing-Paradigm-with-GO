package main

import (
	"fmt"
	"math"
	"time"
)

// isPrime checks if a single number is prime
// a number is prime if it's only divisible by 1 and itself
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	// only check up to the square root of n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false // divisble by something, not prime
		}
	}
	return true

}

// findPrimesSingleThreaded finds all primes up to a limit, one at a time
func findPrimesSingleThread(limit int) []int {
	primes := []int{} //empty slice to store results

	for n := 2; n <= limit; n++ {
		if isPrime(n) {
			primes = append(primes, n) //add to results
		}
	}
	return primes
}

// worker recieves numbers to check from jobs channel
// sends confirmed primes back through results channel
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		if isPrime(n) {
			results <- n //send prime back through results channel
		}
	}
}

func findPrimesParallel(limit int, workerCount int) []int {
	jobs := make(chan int, limit)    //channel to send numbers to workers
	results := make(chan int, limit) //channel to collect primes back

	//launch workerCount goroutines, all listening on jobs channel
	for n := 0; n < workerCount; n++ {
		jobs <- n
	}
	close(jobs) // tells workers no more numbers are coming

	//collect all-results
	primes := []int{}
	for n := 2; n <= limit; n++ {
		result := <-results
		primes = append(primes, result)
	}
	return primes
}

func main() {
	limit := 1000000 //we'll find all primes up to this limit

	fmt.Println("--- Single-threaded prime finder ---")
	fmt.Printf("Finding primes up to %d...\n", limit)

	// starts timer
	start := time.Now()

	//running single-threaded prime finder
	primes := findPrimesSingleThread(limit)

	//stops timer
	elapsed := time.Since(start)

	fmt.Printf("Found %d primes\n", len(primes))
	fmt.Printf("Time taken: %s\n", elapsed)
}
