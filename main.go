package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")
	fmt.Printf("Sum of %d and %d is %d", 5, 6, calculateSum(5, 6))
}

func calculateSum(x int, y int) int {
	return x + y
}
