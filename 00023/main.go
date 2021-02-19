package main //https://leetcode.com/problems/merge-k-sorted-lists/

import "fmt"

// ListNode ..
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var lst = []*ListNode{
		&ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  10,
				Next: nil,
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	var r = mergeKLists(lst)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}

// Time complexity: O(Nlogk)
// Space complexity: O(1)
func mergeKLists(lists []*ListNode) *ListNode { //排除merge2List函数，时间复杂度为: logk
	var size = len(lists)
	if size == 0 {
		return nil
	}

	if size == 1 {
		return lists[0]
	}

	var m = size / 2
	var l1 = mergeKLists(lists[:m])
	var l2 = mergeKLists(lists[m:])

	return merge2List(l1, l2)
}

func merge2List(l1, l2 *ListNode) *ListNode { //时间复杂度: N
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = merge2List(l1.Next, l2)
		return l1
	}

	l2.Next = merge2List(l1, l2.Next)
	return l2
}
