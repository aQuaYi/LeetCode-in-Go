package problem0101

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true 
	}

	return recur(root.Left, root.Right)
}

func recur(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && recur(left.Left, right.Right) && recur(left.Right, right.Left)
}
