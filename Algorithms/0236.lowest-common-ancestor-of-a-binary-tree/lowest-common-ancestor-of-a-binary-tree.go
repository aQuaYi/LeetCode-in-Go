package problem0236

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p.Val == root.Val || q.Val == root.Val {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}

	return r
}
