package problem0110

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func isBalanced(root *TreeNode) bool {
	_, isBalanced := recur(root)
	return isBalanced
}

func recur(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftIsBalanced := recur(root.Left)
	rightDepth, rightIsBalanced := recur(root.Right)

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
