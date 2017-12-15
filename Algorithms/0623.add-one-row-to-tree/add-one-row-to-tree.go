package Problem0623

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		newRoot := newNode(v)
		newRoot.Left = root
		return newRoot
	}
	helper(root, v, d)
	return root
}

func helper(root *TreeNode, v, d int) {
	if d == 2 {
		tl := root.Left
		root.Left = newNode(v)
		root.Left.Left = tl
		tr := root.Right
		root.Right = newNode(v)
		root.Right.Right = tr
		return
	}

	if root.Left != nil {
		helper(root.Left, v, d-1)
	}

	if root.Right != nil {
		helper(root.Right, v, d-1)
	}

}

func newNode(v int) *TreeNode {
	return &TreeNode{
		Val: v,
	}
}
