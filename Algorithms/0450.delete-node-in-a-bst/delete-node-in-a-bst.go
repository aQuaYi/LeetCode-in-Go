package Problem0450

import (
	"fmt"
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		newRight := insert2Left(root.Right, root.Left.Right)
		root.Left.Right = newRight
		return root.Left
	}

	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)

	return root
}

func insert2Left(root, node *TreeNode) *TreeNode {
	if root == nil {
		return node
	}

	if root.Left == nil {
		root.Left = node
		return root
	}

	 insert2Left(root.Left, node)
	return root
}
