package problem1022

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

func sumRootToLeaf(root *TreeNode) int {
	res := 0
	dfs(root, 0, &res)
	return res
}

func dfs(node *TreeNode, num int, res *int) {
	num = num*2 + node.Val
	if node.Left == nil && node.Right == nil {
		*res += num
		return
	}
	if node.Left != nil {
		dfs(node.Left, num, res)
	}
	if node.Right != nil {
		dfs(node.Right, num, res)
	}
}
