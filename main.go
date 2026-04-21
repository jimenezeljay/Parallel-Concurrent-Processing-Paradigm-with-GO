package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// isPrime checks if a single number is prime
// it only checks divisors up to the square root of n,
// which is a standard math optimization to reduce unnecessary checks
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
// this is the baseline single-threaded version - no parallelism
// every number is checked sequentially, one after another
func findPrimesSingleThread(limit int) []int {
	primes := []int{} //empty slice to store results

	for n := 2; n <= limit; n++ {
		if isPrime(n) {
			primes = append(primes, n) //add to results
		}
	}
	return primes
}

// worker is a goroutine func that recieves numbers to check from jobs channel
// sends confirmed primes back through results channel
// multiple workers run simultanesously, each pulling from the same job channel until it's closed
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		if isPrime(n) {
			results <- n //send prime back through results channel
		}
	}
}

// findPrimesParallel distributes the prime-checking workload across multiple goroutines
// it ues two channels: jobs (send numbers to workers) and results (collects primes back)
// a WaitGroup ensures main() waits for ALL workers to finish before collecting results
func findPrimesParallel(limit int, workerCount int) []int {
	jobs := make(chan int, limit)    //channel to send numbers to workers
	results := make(chan int, limit) //channel to collect primes back

	//waitgroup tracks when all workers are done
	var wg sync.WaitGroup

	//launch workerCount goroutines, all listening on jobs channel
	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, results)
		}()
	}

	//send every number from 2 to limit into the job channel
	for n := 2; n <= limit; n++ {
		jobs <- n
	}
	close(jobs) //signal workers no more numbers are coming

	//once all workers finish, close results
	go func() {
		wg.Wait() //wait for all workers to finish
		close(results)
	}()

	//collect all primes sent back through the results channel
	primes := []int{}
	for p := range results {
		primes = append(primes, p)
	}

	return primes
}

func main() {
	limit := 5000000 //find all primes up to 5 million
	workerCount := 8 //number of parallel goroutines to use

	fmt.Println("------------------------------------------")
	fmt.Println("       Parallel Prime Number Finder       ")
	fmt.Println("------------------------------------------")
	fmt.Printf("Finding all primes up to %d...\n\n", limit)

	// --- Single-Threaded Version ---
	fmt.Println("=== Single-Threaded ===")
	start := time.Now()
	primesSingle := findPrimesSingleThread(limit)
	singleTime := time.Since(start)
	fmt.Printf("Found %d primes in %s\n\n", len(primesSingle), singleTime)

	// --- Parallel Version ---
	fmt.Println("=== Parallel (8 Goroutine Workers) ===")
	start = time.Now()
	primesParallel := findPrimesParallel(limit, workerCount)
	parallelTime := time.Since(start)
	fmt.Printf("Found %d primes in %s\n\n", len(primesParallel), parallelTime)

	// --- Performance Comparison ---
	fmt.Println("------------------------------------------")
	fmt.Printf("  Speedup: %.2fx faster with parallel\n",
		float64(singleTime)/float64(parallelTime))
	fmt.Println("------------------------------------------")
}
