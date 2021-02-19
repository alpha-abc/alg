package main // https://leetcode.com/problems/add-two-numbers/

import "fmt"

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int // non-negative, single digit
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r = new(ListNode)
	var head = r

	var k = 0
	for l1 != nil || l2 != nil {
		var n = 0
		if l1 != nil {
			n = n + l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n = n + l2.Val
			l2 = l2.Next
		}

		head.Next = new(ListNode)
		head.Next.Val = (n + k) % 10
		k = (n + k) / 10
		head = head.Next
	}

	if k > 0 {
		head.Next = new(ListNode)
		head.Next.Val = k
	}
	return r.Next
}

func main() {
	var l1 = &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	var l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
	}

	var r = addTwoNumbers(l1, l2)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
