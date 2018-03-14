package problem0124

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val

	var dfs func(*TreeNode) int
	// 返回，以 root 为起点，所有可能路径的 sum 值中的最大值。
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
