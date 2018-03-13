package Problem0783

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = kit.TreeNode

func minDiffInBST(root *TreeNode) int {
	res := 1<<63 - 1
	helper(root, nil, &res)
	return res

}

func helper(root *TreeNode, pre, res *int) {
	if root.Left != nil {
		helper(root.Left, pre, res)
	}

	if pre != nil {
		*res = min(*res, root.Val-*pre)
	}

	pre = &root.Val

	if root.Right != nil {
		helper(root.Right, pre, res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
