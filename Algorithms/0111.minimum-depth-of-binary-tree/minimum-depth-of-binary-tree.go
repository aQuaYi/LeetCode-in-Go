package Problem0111

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

// 注意审题，没有子节点的 node 才是 leaf
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return sonDepth(minDepth(root.Left),minDepth(root.Right ))+1
}

func sonDepth(a, b int) int {
	if a == 0 || b == 0 {
		return a+b 
	}

	if a < b {
		return a
	}

	return b
}
