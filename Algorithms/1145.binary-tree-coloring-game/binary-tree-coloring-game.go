package problem1145

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var left, right int

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := dfs(node.Left), dfs(node.Right)
		if node.Val == x {
			left, right = l, r
		}
		return l + r + 1
	}

	dfs(root)

	up := n - left - right - 1
	n /= 2

	return left > n || right > n || up > n
}

// node-x split tree into 3 parts:
// node-x.Left, node-x.Right and up(node-x parent plus node-x brother)
// second player must takes the biggest part
// second player win the game if his part > n/2
