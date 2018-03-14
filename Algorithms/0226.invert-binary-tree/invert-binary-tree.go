package problem0226

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil ||
		(root.Left == nil && root.Right == nil) {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}
