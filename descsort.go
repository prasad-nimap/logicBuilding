package main

import "fmt"

func main() {
	arr := []int{5, 3, 1, 2, 4}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			temp := arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = temp

			// restart the loop
			i = -1
		}
	}

	fmt.Println(arr)
}
