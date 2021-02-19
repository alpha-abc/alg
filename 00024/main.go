package main // https://leetcode.com/problems/swap-nodes-in-pairs/

import (
	"fmt"
	"time"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * Given 1->2->3->4, you should return the list as 2->1->4->3.

  var arr1 = [ 2, 4]
			   | /|
               |/ |
  var arr2 = [ 1, 3]
*/

// ListNode .
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var refs [2][]*ListNode
	var a, b = 0, 1

	var l = 0
	for i := head; i != nil; i = i.Next {
		if l%2 == 1 {
			refs[a] = append(refs[a], i)
		} else {
			refs[b] = append(refs[b], i)
		}
		l++
	}

	var la = len(refs[a])
	var lb = len(refs[b])

	if l <= 1 {
		return head
	}

	for i := 0; i < la; i++ {
		refs[a][i].Next = refs[b][i]

		if i > 0 {
			refs[b][i-1].Next = refs[a][i]
		}
	}

	if lb > la {
		refs[b][lb-2].Next = refs[b][lb-1]
	}

	refs[b][lb-1].Next = nil
	return refs[a][0]
}

func main() {
	var n = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	var n1 = swapPairs(n)
	for head := n1; n1.Next != nil; head = head.Next {
		fmt.Println(head)
		time.Sleep(1 * time.Second)
	}

}
