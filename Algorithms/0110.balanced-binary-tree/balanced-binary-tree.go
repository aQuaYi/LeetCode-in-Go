package Problem0110

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func isBalanced(root *TreeNode) bool {
	_, isBalanced := cur(root)
	return isBalanced
}

func cur(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftIsBalanced := cur(root.Left)
	rightDepth, rightIsBalanced := cur(root.Right)

	if leftIsBalanced && rightIsBalanced &&
		-1 <= leftDepth-rightDepth  && leftDepth - rightDepth <= 1 {
		return max(leftDepth, rightDepth) + 1, true
	}

	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
