package problem0103

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 出现了新的 level
		if level >= len(res) {
			res = append(res, []int{})
		}
		// 与 102 题相比，level 的奇偶不同，append 会不同。
		if level%2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			temp := make([]int ,  len(res[level])+1)
			temp[0] = root.Val
			copy(temp[1:], res[level])
			res[level] = temp
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
