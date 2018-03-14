package problem0104

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
