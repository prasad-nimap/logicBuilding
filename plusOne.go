package main

import (
	"fmt"
	"strconv"
)

//Input: digits = [1,2,3]
//Output: [1,2,4]

func plusOne(digits []int) []int {
	var nums []int

	lastelement := digits[len(digits)-1]
	//	fmt.Println("Get the last element:", lastelement)

	lastelement = lastelement + 1

	if lastelement > 9 {
		var digit int
		numberStr := strconv.Itoa(lastelement)
		digit, _ = strconv.Atoi(string(numberStr))
		

		nums = digits[0 : len(digits)-1]
		fmt.Println(nums)
		nums = append(nums, digit)

	} else {
		nums = digits[0 : len(digits)-1]
		nums = append(nums, lastelement)
	}

	return nums
}

func main() {
	digits := []int{1, 2, 9}
	fmt.Print(plusOne(digits))
}
