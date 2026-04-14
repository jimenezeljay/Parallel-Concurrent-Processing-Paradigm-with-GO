package main

import "fmt"

func main() {

	//creating a slice
	numbers := []int{10, 20, 30, 40, 50}

	//accessing by index
	fmt.Println("First:", numbers[0])
	fmt.Println("Last:", numbers[4])

	//looping over a slice with range
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	//adding to a slice with append
	numbers = append(numbers, 60)
	fmt.Println("After append:", numbers)
}
