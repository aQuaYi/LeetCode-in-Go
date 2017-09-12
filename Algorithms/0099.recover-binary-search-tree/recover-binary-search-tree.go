package Problem0099

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

// 已经假设 BST 中没有重复的值
// 所以 root != nil
func recoverTree(root *TreeNode) {
	// Golang int 类型的最小值与最大值
	MIN, MAX := -1<<63, 1<<63-1

	left := wrongNodeOfBST(MIN, root.Val, root.Left)
	right := wrongNodeOfBST(root.Val, MAX, root.Right)

	switch {
	case left != nil && right != nil:
		left.Val, right.Val = right.Val, left.Val
	case left != nil:
		left.Val, root.Val = root.Val, left.Val
	case right != nil:
		right.Val, root.Val = root.Val, right.Val
	}
}

// 如果 root.Val 不在 (min, max) 范围内
// root 就是出错的节点
func wrongNodeOfBST(min, max int, root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var res *TreeNode
	if root.Val < min || max < root.Val {
		// 不能立即返回 root，因为
		// 当 root 的父节点和子节点对换后，root 也会被判断为错误节点
		// 例如
		//    preorder = []int{3,2,1}
		//    inorder  = []int{3,2,1}
		//    值为 2 的节点也会被判断为错误节点
		res = root
	}

	node := wrongNodeOfBST(min, root.Val, root.Left)
	if node != nil {
		return node
	}

	node = wrongNodeOfBST(root.Val, max, root.Right)
	if node != nil {
		return node
	}

	// 只有当 root 的左右子树都没有错误节点的时候
	// 才返回 res
	return res
}
