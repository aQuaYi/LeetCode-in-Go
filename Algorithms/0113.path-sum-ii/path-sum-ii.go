package Problem0113

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	path := []int{}

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level, sum int) {
		if node == nil {
			return
		}

		if level >= len(path) {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}
		sum -= node.Val

		if node.Left == nil && node.Right == nil {
			if sum == 0 {
				temp := make([]int, level+1)
				copy(temp, path)
				res = append(res, temp)
			}
		}

		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)
	}

	dfs(root, 0, sum)

	return res
}
