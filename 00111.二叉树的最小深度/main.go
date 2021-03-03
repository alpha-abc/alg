package main // https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

import (
	"container/list"
	"fmt"
)

func main() {
	var root = &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
		},
	}

	fmt.Println(minDepth(root))
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var lst = list.New().Init()

	lst.PushBack(root)
	var depth = 1

	for lst.Len() > 0 {
		var sz = lst.Len()
		for i := 0; i < sz; i++ {
			var e = lst.Front()
			lst.Remove(e)

			var node = e.Value.(*TreeNode)

			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				lst.PushBack(node.Left)
			}

			if node.Right != nil {
				lst.PushBack(node.Right)
			}
		}
		depth++
	}
	return depth
}
