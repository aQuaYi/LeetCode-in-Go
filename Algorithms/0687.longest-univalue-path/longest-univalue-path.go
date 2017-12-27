package Problem0687

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	helper(root, &maxLen)
	return maxLen
}

func helper(root *TreeNode, maxLen *int) {
	if root == nil {
		return
	}

	*maxLen = max(
		*maxLen,
		count(root.Left, root.Val)+count(root.Right, root.Val),
	)

	helper(root.Left, maxLen)
	helper(root.Right, maxLen)
}

func count(root *TreeNode, val int) int {
	if root == nil || root.Val != val {
		return 0
	}
	return 1 + max(count(root.Left, val), count(root.Right, val))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
