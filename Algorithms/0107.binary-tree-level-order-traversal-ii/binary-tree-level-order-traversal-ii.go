package Problem0107

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
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
			temp := make([][]int, len(res )+1)
			temp[0]= []int{}
			copy(temp[1:], res )
			res = temp 
		}
		n := len(res )
		res[n-level-1] = append(res[n-level-1], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res 
}
