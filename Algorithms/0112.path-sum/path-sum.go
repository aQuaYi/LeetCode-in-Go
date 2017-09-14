package Problem0112

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return dfs(root, sum)
}

func dfs(root *TreeNode, sum int) bool {
	if root == nil {
		if sum == 0 {
			return true
		}
		return false
	}

	temp := sum - root.Val
	switch {
	case root.Left == nil:
		return dfs(root.Right, temp)
	case root.Right == nil:
		return dfs(root.Left, temp)
	default:
		return dfs(root.Left, temp) || dfs(root.Right, temp)
	}
}
