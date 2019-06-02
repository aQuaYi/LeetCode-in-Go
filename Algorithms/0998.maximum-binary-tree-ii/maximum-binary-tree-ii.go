package problem0998

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val > root.Val {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}

	root.Right = insertIntoMaxTree(root.Right, val)

	return root
}
