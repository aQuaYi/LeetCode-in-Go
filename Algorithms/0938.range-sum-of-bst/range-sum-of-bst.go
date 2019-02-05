package problem0938

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is Definition for a binary tree node.
type TreeNode = kit.TreeNode

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	sum := 0

	switch {
	case root.Val < L:
		sum = rangeSumBST(root.Right, L, R)
	case R < root.Val:
		sum = rangeSumBST(root.Left, L, R)
	default:
		sum = root.Val
		sum += rangeSumBST(root.Left, L, R)
		sum += rangeSumBST(root.Right, L, R)
	}

	return sum
}
