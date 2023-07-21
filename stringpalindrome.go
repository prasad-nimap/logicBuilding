package main

import "fmt"

func main() {
	str := "bob"
	temp := str

	var reversedstring string

	for _, val := range str {
		reversedstring = string(val) + reversedstring
	}

	if reversedstring == temp {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
