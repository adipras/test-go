package main

import (
	"fmt"
)

func main() {
	arrInt := []int{7, 4, 8, 2, 9}

	fmt.Println(CountBuildings(arrInt, len(arrInt)))
}

func CountBuildings(arr []int, n int) int {
	// Initialize result (Note that first building
	// always sees sunlight)
	var count int = 1

	// Start traversing element
	var curr_max int = arr[0]
	for i := 1; i < n; i++ {
		// If curr_element is maximum
		// or current element is
		// equal, update maximum and increment count
		if arr[i] > curr_max || arr[i] == curr_max {
			count++
			curr_max = arr[i]
		}
	}

	return count
}
