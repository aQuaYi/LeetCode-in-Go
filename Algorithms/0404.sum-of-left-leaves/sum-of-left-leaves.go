package Problem0404

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	if root==nil ||
	(root.Left == nil && root.Right == nil ) {
		return 0
	}

	if root.Left.Left == nil && root.Left.Right==nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	return sumOfLeftLeaves(root.Left)+ sumOfLeftLeaves(root.Right)
}
