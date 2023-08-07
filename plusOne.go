package main

import (
	"fmt"
)

//Input: digits = [1,2,3]
//Output: [1,2,4]

func plusOne(digits []int) []int {
	var nums []int

	lastelement := digits[len(digits)-1]
	//	fmt.Println("Get the last element:", lastelement)

	lastelement = lastelement + 1

	if lastelement > 9 { //10
		var newarr []int
		for lastelement > 0 {
			remainder := lastelement % 10 // 0
			lastelement /= 10
			newarr = append(newarr, remainder)
			//			fmt.Println("new arr:", newarr)
		}
		//		fmt.Println("new arr outer:", newarr)

		//1 2 9
		//		fmt.Println("nums ", nums)

		nums = digits[0 : len(digits)-1] // 1 2
		//		fmt.Println("new arr length:", len(newarr))

		for i := len(newarr) - 1; i > -1; i-- {
			//			fmt.Println("Entering the loop")
			nums = append(nums, newarr[i]) // 1 2 1 0
			//			fmt.Println("append nums", nums)

		}

	} else {
		nums = digits[0 : len(digits)-1]
		nums = append(nums, lastelement)
	}

	return nums
}

func main() {
	digits := []int{1, 2, 10}
	fmt.Print(plusOne(digits))
}
