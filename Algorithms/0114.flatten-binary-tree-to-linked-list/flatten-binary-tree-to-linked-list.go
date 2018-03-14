package problem0114

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func flatten(root *TreeNode) {
	recur(root)
}

// recur 会 flatten root，并返回 flatten 后的 leaf 节点
func recur(root *TreeNode) *TreeNode {
	if root == nil ||
		(root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil {
		return recur(root.Right)
	}

	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return recur(root.Right)
	}

	res := recur(root.Right)
	recur(root.Left).Right = root.Right
	root.Right = root.Left
	// 不要忘记封闭 left 节点
	root.Left = nil

	return res
}
