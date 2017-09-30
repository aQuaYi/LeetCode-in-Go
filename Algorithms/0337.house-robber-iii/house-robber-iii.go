package Problem0337

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func rob(root *TreeNode) int {
	res := robSub(root)
	return max(res[0], res[1])
}

func robSub(root *TreeNode) []int {
	if root == nil {
		return make([]int, 2)
	}

	left := robSub(root.Left)
	right := robSub(root.Right)

	res := make([]int, 2)
	res[0] = max(left[0], left[1]) + max(right[0], right[1])
	res[1] = root.Val + left[0] + right[0]

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
