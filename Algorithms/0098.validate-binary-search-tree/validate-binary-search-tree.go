package Problem0098

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}

	if root.Right != nil && root.Val >= root.Right.Val {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}
