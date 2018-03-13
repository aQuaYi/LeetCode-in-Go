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

const MAXINT = 1<<63 - 1

func minDiffInBST(root *TreeNode) int {
	res := MAXINT

	if root.Left != nil {
		res = min(res, root.Val-root.Left.Val)
		res = min(res, minDiffInBST(root.Left))
	}

	if root.Right != nil {
		res = min(res, root.Right.Val-root.Val)
		res = min(res, minDiffInBST(root.Right))
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
