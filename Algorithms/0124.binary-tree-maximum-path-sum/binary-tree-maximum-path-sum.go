package Problem0124

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val

	// 返回，从 root 出发，包含 root 在内的所有可能路径的最大的 sum 值
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		sum := left + root.Val + right
		if maxSum < sum {
			maxSum = sum
		}

		return max(left, right) + root.Val
	}

	dfs(root)

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
