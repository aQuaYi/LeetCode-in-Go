package problem0623

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	switch d {
	case 1:
		newRoot := &TreeNode{Val: v, Left: root}
		return newRoot
	case 2:
		left := &TreeNode{Val: v, Left: root.Left}
		right := &TreeNode{Val: v, Right: root.Right}
		root.Left = left
		root.Right = right
	default:
		if root.Left != nil {
			root.Left = addOneRow(root.Left, v, d-1)
		}
		if root.Right != nil {
			root.Right = addOneRow(root.Right, v, d-1)
		}
	}

	return root
}
