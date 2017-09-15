package Problem0124

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return -1<<31 
	}

	sum := root.Val +
		max(0, maxSince(root.Left)) +
		max(0, maxSince(root.Right))

	return max(sum,
		max(maxPathSum(root.Left), maxPathSum(root.Right)))
}

// 返回，从 root 出发，包含 root 在内的所有可能路径的最大的 sum 值
func maxSince(root *TreeNode) int {
	res := -1 << 31

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}

		sum += root.Val
		if res < sum {
			res = sum
		}

		dfs(root.Left, sum)
		dfs(root.Right, sum)
	}

	dfs(root, 0)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
