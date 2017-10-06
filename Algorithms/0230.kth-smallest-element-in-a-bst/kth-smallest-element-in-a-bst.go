package Problem0230

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	leftSize := getSize(root.Left)
	switch {
	case k <= leftSize:
		return kthSmallest(root.Left, k)
	case leftSize+1 < k:
		return kthSmallest(root.Right, k-leftSize-1)
	default:
		return root.Val
	}
}


func getSize(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + getSize(root.Left) + getSize(root.Right)
}
