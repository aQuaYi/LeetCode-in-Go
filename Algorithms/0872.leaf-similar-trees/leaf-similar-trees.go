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
	return equal(leafValues(root1), leafValues(root2))
}

func leafValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	return append(leafValues(root.Left), leafValues(root.Right)...)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
