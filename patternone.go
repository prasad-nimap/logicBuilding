package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print(j, " ")
		}
		fmt.Println("")
	}
}
/*
5 4 3 2 1 
5 4 3 2
5 4 3
5 4
5
*/