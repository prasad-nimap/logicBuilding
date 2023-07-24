package main

import "fmt"

func main() {
	rows := 6

	for i := 1; i <= rows; i++ {
		for j := i; j <= rows; j++ {
			 fmt.Print("*", " ")
		}
		fmt.Println("")
	}
}
