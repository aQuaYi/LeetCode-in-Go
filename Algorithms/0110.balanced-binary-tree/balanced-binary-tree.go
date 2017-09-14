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

	leftDeepth, leftIsBalanced := cur(root.Left)
	rightDeepth, rightIsBalanced := cur(root.Right)

	if leftIsBalanced && rightIsBalanced &&
		-1 <= leftDeepth-rightDeepth  && leftDeepth - rightDeepth <= 1 {
		return max(leftDeepth, rightDeepth) + 1, true
	}

	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
