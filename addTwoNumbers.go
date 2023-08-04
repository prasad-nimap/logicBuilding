package main

import "fmt"

/* // Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
} 
type ListNode struct {
	Data int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	carry := 0

	current1 := l1.Data
	current2 := l2.Data

	for l1 != nil || l2 != nil {
		// var l1Val, l2Val int
		if l1 != nil {
			current1 = l1.Data
			l1 = l1.Next
		}
		if l2 != nil {
			current2 = l2.Data
			l2 = l2.Next
		}

		sum := current1 + current2 + carry
		carry = sum / 10

		current.Next = &ListNode{Data: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Data: carry}
	}

	return head.Next
}
*/
func main() {
	x := 123
	reverseNumber := 0

	for x > 0 {
		reaminder := x % 10
		reverseNumber = (reverseNumber * 10) + reaminder
		reverseNumber /= 10
	}

	fmt.Println(reverseNumber)
}
