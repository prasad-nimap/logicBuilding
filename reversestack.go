package main

import "fmt"

type stack []int

func (s *stack) push(nums int) {
	*s = append(*s, nums)
}

func (s *stack) pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	index := len(*s) - 1
	top := (*s)[index]
	*s = (*s)[:index]
	return top, true
}

func reverse(nums []int) []int {
	var s stack
	reversed := make([]int, 0, len(nums))
	for _, val := range nums {
		s.push(val)
	}

	for {
		if val, ok := s.pop(); ok {
			reversed = append(reversed, val)
		} else {
			break
		}
	}
	return reversed
}

func main() {
	nums := []int{6, 5, 4, 3, 2, 1}
	fmt.Println(reverse(nums))
}
