package problem1026

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func dfs(node *TreeNode, Max, Min int) int {
	if node == nil {
		return Max - Min
		// Max 和 Min 都是 node 的父节点的值
		// 所以
		// Max 和 Min 所在的两个节点，肯定是父子关系
	}

	Max = max(Max, node.Val)
	Min = min(Min, node.Val)

	return max(dfs(node.Left, Max, Min), dfs(node.Right, Max, Min))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
