package problem0965

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined in kit
type TreeNode = kit.TreeNode

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if (root.Left != nil && root.Val != root.Left.Val) ||
		(root.Right != nil && root.Val != root.Right.Val) {
		return false
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
