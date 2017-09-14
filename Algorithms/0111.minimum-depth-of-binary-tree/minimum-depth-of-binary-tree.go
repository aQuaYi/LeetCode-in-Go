package Problem0111

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 题目很奇葩的规定
	//     inorder[1]     的最短深度为 1
	//     inorder[1, 2]  的最短深度为 2
	if root.Left == nil && root.Right == nil {
		return 1
	}

	return max(2, min(minDepth(root.Left),minDepth(root.Right ))+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
