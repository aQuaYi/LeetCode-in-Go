package Problem0437

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	res := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum -= node.Val
		if sum == 0 {
			res++
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	dfs(root, sum)

	return res +
		pathSum(root.Left, sum) +
		pathSum(root.Right, sum)
}
