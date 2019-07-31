package problem1123

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	res, _ := dfs(root)
	return res
}

func dfs(node *TreeNode) (*TreeNode, int) {
	if node == nil {
		return nil, 0
	}

	l, ld := dfs(node.Left)
	r, rd := dfs(node.Right)

	if ld > rd {
		return l, ld + 1
	}

	if rd > ld {
		return r, rd + 1
	}

	return node, ld + 1
}
