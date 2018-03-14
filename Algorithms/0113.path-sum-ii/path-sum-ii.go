package problem0113

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level, sum int) {
		if node == nil {
			return
		}

		// 根据 level 来更新 path
		if level >= len(path) {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}
		sum -= node.Val

		// 到达 leaf
		// 并且，此条路径符合要求
		if node.Left == nil && node.Right == nil && sum == 0 {
			temp := make([]int, level+1)
			// copy 不会复制 path[len(temp):]，如果 len(path) > len(temp) 的话
			copy(temp, path)
			res = append(res, temp)
		}

		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)
	}

	dfs(root, 0, sum)

	return res
}
