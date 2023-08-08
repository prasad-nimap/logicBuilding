package main

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := (len(nums) - 1)

	// fmt.Println(nums[mid], low, high)
	for low <= high {
		//mid := (low + (high - low)) / 2
		mid := (low + high) / 2

		if nums[mid] < target { // <- move towards right
			low = mid + 1
		} else if nums[mid] > target { // <- move towards left
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	nums := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23
	fmt.Print("Target element is present on index:",search(nums, target))
}
