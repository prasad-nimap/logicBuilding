/*
// TODO:
* * * * *
* * * * *
* * * * *
*/

package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print("*", " ")
		}

		fmt.Print("\n")
	}
}
