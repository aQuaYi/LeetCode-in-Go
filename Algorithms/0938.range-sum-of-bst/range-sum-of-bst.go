package problem0938

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is Definition for a binary tree node.
type TreeNode = kit.TreeNode

func rangeSumBST(root *TreeNode, L int, R int) int {
	res := 0
	recur(root, L, R, &res)
	return res
}

func recur(node *TreeNode, L, R int, res *int) {
	if node == nil {
		return
	}

	if L <= node.Val && node.Val <= R {
		*res += node.Val
	}

	recur(node.Left, L, R, res)
	recur(node.Right, L, R, res)

	return
}
