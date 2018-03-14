package problem0337

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	// a 抢劫 root 节点时的最大值
	// b 不抢劫 root 节点时的最大值
	dfs = func(root *TreeNode) (a, b int) {
		if root == nil {
			return 0, 0
		}
		la, lb := dfs(root.Left)
		ra, rb := dfs(root.Right)
		a = root.Val + lb + rb
		b = max(la, lb) + max(ra, rb)
		return a, b
	}

	a, b := dfs(root)
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
