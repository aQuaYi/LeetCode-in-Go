package Problem0112

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return cur(root, sum)
}

func cur(root *TreeNode, sum int) bool {
	if root == nil {
		if sum == 0 {
			return true
		}
		return false
	}

	if root.Val > sum {
		return false
	}

	return cur(root.Left, sum-root.Val) || cur(root.Right, sum-root.Val)
}
