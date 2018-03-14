package problem0617

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	root := &TreeNode{Val: t1.Val + t2.Val}

	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)

	return root
}
