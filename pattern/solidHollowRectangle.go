package main

import "fmt"

func main() {
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows; j++ {
			fmt.Print("*", " ")
			/*		if i == 1 || j == i || j == rows {
						fmt.Print("*", " ")
					} else {
						fmt.Print(" ")
					}*/
		}
		fmt.Println("")
	}
}
