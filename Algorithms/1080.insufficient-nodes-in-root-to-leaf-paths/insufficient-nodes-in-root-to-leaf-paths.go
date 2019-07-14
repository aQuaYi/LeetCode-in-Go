package problem1080

import (
	"math"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if dfs(root, 0, limit) < limit {
		return nil
	}
	return root
}

func dfs(node *TreeNode, pre, limit int) int {
	if node.Left == nil && node.Right == nil {
		return node.Val
	}

	pre += node.Val

	l, r := math.MinInt64, math.MinInt64

	if node.Left != nil {
		l = dfs(node.Left, pre, limit)
		if pre+l < limit {
			node.Left = nil
		}
	}

	if node.Right != nil {
		r = dfs(node.Right, pre, limit)
		if pre+r < limit {
			node.Right = nil
		}
	}

	return max(l, r) + node.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
