package problem0230

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	leftSize := getSize(root.Left)
	switch {
	case k <= leftSize:
		// 答案存在于 root.Left 中
		return kthSmallest(root.Left, k)
	case leftSize+1 < k:
		// 答案存在于 root.Right 中
		return kthSmallest(root.Right, k-leftSize-1)
	default:
		// 答案是 root.Val
		return root.Val
	}
}

// 获取 root 树的节点数量
func getSize(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getSize(root.Left) + getSize(root.Right)
}
