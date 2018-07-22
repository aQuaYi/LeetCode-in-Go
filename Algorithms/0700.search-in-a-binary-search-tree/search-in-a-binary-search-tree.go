package problem0700

import "github.com/aQuaYi/LeetCode-in-Go/kit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode 是预定义的类型
type TreeNode = kit.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case root.Val < val:
		return searchBST(root.Right, val)
	case val < root.Val:
		return searchBST(root.Left, val)
	default:
		return root
	}
}
