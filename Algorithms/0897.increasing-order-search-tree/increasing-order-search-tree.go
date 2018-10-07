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
	if root == nil {
		return nil, nil
	}

	if root.Left == nil && root.Right == nil {
		return root, root
	}

	if root.Left != nil {
		leftRoot, leftRight := helper(root.Left)
		leftRight.Right = root
		if root.Right != nil {
			rightRoot, rightRight := helper(root.Right)
			root.Right = rightRoot
			return leftRoot, rightRight
		} else {
			return leftRoot, root
		}
	}
	rightRoot, rightRight := helper(root.Right)
	root.Right = rightRoot
	return root, rightRight
}
