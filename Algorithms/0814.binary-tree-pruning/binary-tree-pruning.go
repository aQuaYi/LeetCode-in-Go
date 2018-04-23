package problem0814

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * 	return nil
}
*/

// TreeNode 是 tree 的 node
type TreeNode = kit.TreeNode

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Val == 0 &&
		root.Left == nil &&
		root.Right == nil {
		return nil
	}
	return root
}
