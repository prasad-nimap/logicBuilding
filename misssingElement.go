// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 5}
	//fmt.Println(len(nums))
	//fmt.Println(len(nums) - 1)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 == nums[i+1] {
			continue
		} else {
			fmt.Println(nums[i] + 1)
			//			fmt.Println(i + 1)
		}
	}
}
