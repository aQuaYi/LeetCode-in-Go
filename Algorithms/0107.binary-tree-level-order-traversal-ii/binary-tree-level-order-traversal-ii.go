package problem0107

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode =  kit.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的 level
		if  level >= len(res)   {
			res = append([][]int{[]int{}}, res... )
		}
		n := len(res )
		res[n-level-1] = append(res[n-level-1], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res 
}
