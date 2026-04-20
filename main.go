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
	jobs := make(chan int, limit) //channel to send numbers to workers
	results := make(chan int, limit) //channel to collect primes back

	//launch workerCount goroutines, all listening on jobs channel
	for w := 0; w < workerCount; w++ {
		go worker(jobs, results)
	}

	//send all numbers to jobs channel
	for n := 2; n <= limit; n++ {
		jobs <- n
	}
	close(jobs) // no more jobs to send

	//collect results - use a counter instead of fixed loop
	primes := []int{}
	for i := 0; i < workerCount; i++ {
		go func() {
			for p := range results {
				primes = append(primes, p)
			}
		}()
	}

	//wait for all workers to finish
	time.Sleep(2 * time.Second)

	return primes
}

func main() {
	limit := 1000000 //we'll find all primes up to this limit
	workerCount := 8 //number of parallel workers to use

	fmt.Println("----------------------------")
	fmt.Println("Parallel Prime Number Finder")
	fmt.Println("----------------------------")

	//single-threaded version
	fmt.Println("--- Single-threaded prime finder ---")
	start := time.Now()
	primesSingle := findPrimesSingleThread(limit)
	singleTime := time.Since(start)
	fmt.Printf("Found %d primes in %s\n\n", len(primesSingle), singleTime)

	//parallel version
	fmt.Println("--- Parallel (8 workers) ---")
	start = time.Now()
	primesParallel := findPrimesParallel(limit, workerCount)
	parallelTime := time.Since(start)
	fmt.Printf("Found %d primes in %s\n\n", len(primesParallel), parallelTime)

	//comparison
	fmt.Println("------------------------------------------")
	fmt.Printf("Speedup: %.2fx faster with parallel\n",
		float64(singleTime)/float64(parallelTime))
	fmt.Println("------------------------------------------")
}
