package problem0222

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is predefined in kit
type TreeNode = kit.TreeNode

func countNodes(root *TreeNode) int {
	count := 0
	traverse(root, &count)
	return count
}

func traverse(n *TreeNode, count *int) {
	if n == nil {
		return
	}

	*count++

	traverse(n.Left, count)
	traverse(n.Right, count)
}
