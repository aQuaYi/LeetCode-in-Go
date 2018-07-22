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
	if root == nil || root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
}
