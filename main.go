package main

import "fmt"

func double(number int, ch chan int) {
	result := number * 2
	ch <- result // Send result back through the channel
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Create a channel that carries ints
	ch := make(chan int)

	// Launch one goroutine per number
	for _, num := range numbers {
		go double(num, ch)
	}

	// Collect ALL results back from the channel
	for i := 0; i < len(numbers); i++ {
		result := <-ch
		fmt.Println("Got result:", result)
	}

	fmt.Println("All done!")
}
