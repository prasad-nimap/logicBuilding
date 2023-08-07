package main

import (
	"fmt"
	"strconv"
)

func divisibilityArray(word string, m int) []int {
	var result []int
	wordtoint, _ := strconv.Atoi(word)

	for i := 0; i < len(wordtoint); i++ {

	}
	return result
}

func main() {
	word := "998244353"
	m := 12

	fmt.Println(divisibilityArray(word, m))

}
