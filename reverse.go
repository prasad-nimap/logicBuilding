package main

import (
	"fmt"
)

func main() {
	num := 123
	reverse := 0

	for num != 0 {
		remainder := num % 10
		reverse = (reverse * 10) + remainder
		num /= 10
	}

	fmt.Println(reverse)
}
