package Problem0337

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	robRoot := root.Val
	if root.Left != nil {
		robRoot += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		robRoot += rob(root.Right.Left) + rob(root.Right.Right)
	}

	noRoot := rob(root.Left)+rob(root.Right)
	return max(robRoot, noRoot )
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
