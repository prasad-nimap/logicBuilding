package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				fmt.Println(nums[i])
			}
		}
	}
}
