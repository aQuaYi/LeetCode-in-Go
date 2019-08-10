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
	if x == root.Val {
		return count(root.Left, x, true) != count(root.Right, x, true)
	}
	return count(root, x, false) <= n/2
}

// count node and its children
func count(node *TreeNode, x int, ok bool) int {
	if node == nil {
		return 0
	}
	if ok || node.Val == x {
		return 1 + count(node.Left, x, true) + count(node.Right, x, true)
	}
	return count(node.Left, x, ok) + count(node.Right, x, ok)
}
