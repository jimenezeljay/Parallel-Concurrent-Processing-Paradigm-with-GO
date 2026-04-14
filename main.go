package main

import "fmt"

func main() {

	//basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Counter:", i)
	}

	//for loop acting as a while loop
	x := 0
	for x < 3 {
		fmt.Println("x is:", x)
		x++
	}
}
