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
	node := find(root, x)
	l, r := count(node.Left), count(node.Right)
	p := n - 1 - l - r
	n /= 2
	return l > n || r > n || p > n
}

func find(node *TreeNode, x int) *TreeNode {
	if node == nil || node.Val == x {
		return node
	}
	l, r := find(node.Left, x), find(node.Right, x)
	if l == nil {
		return r
	}
	return l
}

// count node and its children
func count(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + count(node.Left) + count(node.Right)
}
