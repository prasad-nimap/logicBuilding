package main

import "fmt"

func main() {
	num := 121
	reverse := 0
	
	//	sort the initial number in temp variable
	temp := num

	for num != 0 {
		reaminder := num % 10
		reverse = (reverse * 10) + reaminder
		num /= 10
	}

	fmt.Println(reverse)

	if reverse == temp {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
