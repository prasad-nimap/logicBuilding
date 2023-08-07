// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func reverse(word string) []string {
	var result []string

	for i := len(word) - 1; i >= 0; i-- {
		result = append(result, string(word[i]))
	}

	return result
}

func main() {
	//words := []string{"Paddy", "Prasad", "Junghare"}
	words := []string{"Paddy"}

	for _, val := range words {
		fmt.Println(val)
		fmt.Print(reverse(val))
	}
}
