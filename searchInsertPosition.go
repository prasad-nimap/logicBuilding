package main

import (
	"fmt"
	"sort"
)

// Input: nums = [1,3,5,6], target = 5
// Output: 2

// Input: nums = [1,3,5,6], target = 7
// Output: 4

func searchPosition(nums []int, target int) int {
	for key, val := range nums {
		if val == target {
			return key
		} else if target != val {
			nums = append(nums, target)
			sort.Ints(nums)
			for key, val := range nums {
				if val == target {
					return key
				}
			}
		}
		/* if val == target {
		   return key
		 } */
	}
	return -1
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7

	fmt.Print(searchPosition(nums, target))
}
