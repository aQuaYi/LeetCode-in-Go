package problem0111

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

// 注意审题，没有子节点的 node 才是 leaf
func minDepth(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
