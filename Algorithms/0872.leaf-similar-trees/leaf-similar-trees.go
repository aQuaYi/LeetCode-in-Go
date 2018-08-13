package problem0872

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode definition for a binary tree node.
//  type TreeNode struct {
//      Val int
//      Left *TreeNode
//      Right *TreeNode
//  }
type TreeNode = kit.TreeNode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	a1 := [100]int{}
	a2 := [100]int{}
	search(root1, 0, &a1)
	search(root2, 0, &a2)
	return a1 == a2
}

func search(root *TreeNode, i int, ap *[100]int) int {
	if root == nil {
		return i
	}
	if root.Left == nil && root.Right == nil {
		(*ap)[i] = root.Val
		return i + 1
	}

	i = search(root.Left, i, ap)
	return search(root.Right, i, ap)
}
