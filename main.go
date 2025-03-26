package main

import (
	"fmt"
)

// The Go Programming Language Book

func main() {

	number := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Array %v\n", number)
	fmt.Print("Array ", number[len(number)-1])

	// numbers := [...] int {10, 20, 30, 40, 50}

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Printf("Matrix %v\n", matrix)

	// allN := number[:]
	// first := number[0:3]

	fruits := []string{"apple", "banana", "orange"}
	fmt.Printf("Fruits %v\n", fruits)

	fruits = append(fruits, "grape")

}

//func whit smallcase is private(not exported), with uppercase is public can be exported

func add(a int, b int) (int, int) {
	return a + b, a * b
}
