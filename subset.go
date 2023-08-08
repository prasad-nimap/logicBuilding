package main

import (
	"fmt"
	"sort"
)

func naiveApproach(arr1, arr2 []int) bool {
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr1); j++ {
			if arr2[i] == arr1[j] {
				break
			} else if j == len(arr1) {
				return false
			} else {
				return true
			}
		}
	}

	return false
}

func usingBinarySearchAndSorting(arr1, arr2 []int) bool {
	sort.Ints(arr1)

	var low, high int = 0, (len(arr1) - 1)

	for i := 0; i < (len(arr2) - 1); i++ {
		for low <= high {
			mid := (low + high) / 2
			if arr1[mid] == arr2[i] {
				return true
			} else if arr1[mid] < arr2[i] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}

func main() {
	// arr1 := []int{11, 1, 13, 21, 3, 7}
	arr1 := []int{3, 5, 7, 12, 1, 9, 10, 0, 2}
	// arr2 := []int{11, 3, 7, 1}
	arr2 := []int{6, 3, 8}

	// Naive approach
	// fmt.Print(naiveApproach(arr1, arr2))

	// using sorting and binary search
	fmt.Println(usingBinarySearchAndSorting(arr1, arr2))

}
