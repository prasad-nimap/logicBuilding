package main

type Stack []rune

// push method
func (s *Stack) Push(char rune) {
	*s = append(*s, char)
}

// pop method
func (s *Stack) Pop() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	index := len(*s) - 1 // key of last element
	char := (*s)[index]  // value of last element
	*s = (*s)[:index]    // remove the last element from slice

	return char, true // return char
}

// Reverse
func reversestring(input string) string {
	stack := Stack{}
	var reversed string

	for _, val := range input {
		stack.Push(val)
	}

	for {
		if char, ok := stack.Pop(); ok {
			reversed += string(char)
		} else {
			break
		}
	}

	return reversed
}

func main() {
	word := "paddy"
	println(reversestring(word))
}
