package problem0979

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func distributeCoins(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	l, r := dfs(node.Left, res), dfs(node.Right, res)
	*res += abs(l) + abs(r)
	return l + r + node.Val - 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
