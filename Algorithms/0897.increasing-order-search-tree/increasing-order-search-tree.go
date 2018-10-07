package problem0897

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	res, _ := helper(root)
	return res
}

func helper(root *TreeNode) (*TreeNode, *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return root, root
	}

	left, right := root.Left, root.Right
	root.Left, root.Right = nil, nil

	if left != nil {
		leftRoot, leftRight := helper(left)
		leftRight.Right = root
		if right != nil {
			rightRoot, rightRight := helper(right)
			root.Right = rightRoot
			return leftRoot, rightRight
		}
		return leftRoot, root
	}

	rightRoot, rightRight := helper(right)
	root.Right = rightRoot
	return root, rightRight
}
