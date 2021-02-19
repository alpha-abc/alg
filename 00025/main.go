package main // https://leetcode.com/problems/reverse-nodes-in-k-group/
/**
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
k is a positive integer and is less than or equal to the length of the linked list.
If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:
Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5
For k = 3, you should return: 3->2->1->4->5

Note:
	Only constant extra memory is allowed.
	You may not alter the values in the list's nodes, only nodes itself may be changed.
*/
import "fmt"

func main() {
	var l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	printList(l)

	printList(reverseKGroup(l, 2))

}

func printList(l *ListNode) {
	for {
		if l == nil {
			break
		}
		fmt.Print(l.Val, " -> ")
		l = l.Next
	}
	fmt.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode .
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	var l = 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	if l < k {
		return head
	}

	var dummy *ListNode = new(ListNode)
	dummy.Next = head

	var p1 *ListNode = dummy
	var p2 *ListNode = dummy.Next

	for i := 0; i < l/k; i++ {
		for j := 0; j < k-1; j++ {
			var tmp *ListNode = p2.Next
			p2.Next = p2.Next.Next
			tmp.Next = p1.Next

			/*if i == 0 then: dummy.Next = tmp*/
			p1.Next = tmp
		}

		p1 = p2
		p2 = p2.Next
	}

	return dummy.Next
}
