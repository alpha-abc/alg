package main //https://leetcode.com/problems/remove-nth-node-from-end-of-list/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	var i, j = 0, 0
	var l1, l2 = head, head
	for {
		l1 = l1.Next
		if l1 == nil {
			break
		}

		if i >= n {
			l2 = l2.Next
			j++
		}
		i++
	}

	var lstSize = i + 1
	if lstSize-1 >= j+n {
		l2.Next = l2.Next.Next
		return head
	}

	return head.Next
}

func main() {
	var s = 2
	var n = 4

	var list *ListNode
	var l *ListNode
	for i := 1; i <= s; i++ {
		if i == 1 {
			list = &ListNode{
				Val:  i,
				Next: nil,
			}

			l = list
			continue
		}

		l.Next = &ListNode{
			Val:  i,
			Next: nil,
		}

		l = l.Next
	}

	var p = list
	fmt.Print("[")
	for {
		fmt.Print(p.Val, " ")
		p = p.Next
		if p == nil {
			break
		}
	}
	fmt.Println("]", n)

	var curr = removeNthFromEnd(list, n)
	fmt.Println(".....")
	fmt.Print("[")
	for {
		if curr == nil {
			break
		}
		fmt.Print(curr.Val, " ")
		if curr.Next == nil {
			break
		}
		curr = curr.Next
	}

	fmt.Println("]")
}
