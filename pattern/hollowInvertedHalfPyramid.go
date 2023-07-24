package main

import "fmt"

func main() {
	rows := 5

	for i := 1; i <= rows; i++ {
		for j := i; j <= rows; j++ {
			if j == i || j == rows || i == 1 {
				fmt.Print("*", " ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
