package problem0236

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if isAncestor(p, q) {
		return p
	}
	if isAncestor(q, p) {
		return q
	}
	return helper(root, p, q)
}

func helper(root, p, q *TreeNode) *TreeNode {
	if isAncestor(root.Left, p) &&
		isAncestor(root.Left, q) {
		return helper(root.Left, p, q)
	}
	if isAncestor(root.Right, p) &&
		isAncestor(root.Right, q) {
		return helper(root.Right, p, q)
	}
	return root
}

func isAncestor(p, c *TreeNode) bool {
	if p == nil {
		return false
	}
	if p == c {
		return true
	}
	return isAncestor(p.Left, c) || isAncestor(p.Right, c)
}
