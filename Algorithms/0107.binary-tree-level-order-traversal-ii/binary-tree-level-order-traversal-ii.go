package Problem0107

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode =  kit.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	temp := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的 level
		if  level >= len(temp)   {
			temp = append(temp, []int{})
		}
		temp[level] = append(temp[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	n  := len(temp )
	res := make([][]int , n )
	for i:=0; i<n ; i++{
		res[i] = temp[n-i-1]
	}

	return res 
}
