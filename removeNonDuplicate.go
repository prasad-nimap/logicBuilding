package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 1, 1}

	// Create a map to store the count of occurrences for each element
	countMap := make(map[int]int)

	// Iterate through the array to populate the countMap
	for _, num := range nums {
		countMap[num]++
	}

	// Find the non-duplicate element and print it
	var result int
	for num, count := range countMap {
		if count == 1 {
			result = num
			break // We've found the non-duplicate element, so exit the loop
		}
	}

	fmt.Println("Non-duplicate element:", result)
}
