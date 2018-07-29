package problem0865

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode 是预定义的数据结构
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = kit.TreeNode

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	res, _ := walk(root)
	return res
}

// walk 就是 subtreeWithAllDeepest
// 只是多返回树的深度结果
func walk(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	lst, l := walk(root.Left)
	rst, r := walk(root.Right)

	switch {
	case l > r:
		return lst, l + 1
	case l < r:
		return rst, r + 1
	default:
		return root, r + 1
	}
}
