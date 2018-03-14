package problem0653

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func findTarget(root *TreeNode, k int) bool {
	return helper(root, root, k)
}

func helper(root, searchRoot *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	return (root.Val*2 != k && findNode(searchRoot, k-root.Val)) ||
		helper(root.Left, searchRoot, k) ||
		helper(root.Right, searchRoot, k)
}

func findNode(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	if root.Val < target {
		return findNode(root.Right, target)
	}
	return findNode(root.Left, target)
}
