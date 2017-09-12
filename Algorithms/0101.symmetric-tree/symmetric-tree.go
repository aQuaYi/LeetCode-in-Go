package Problem0101

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true 
	}

	return cur(root.Left, root.Right)
}

func cur(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && cur(left.Left, right.Right) && cur(left.Right, right.Left)
}
