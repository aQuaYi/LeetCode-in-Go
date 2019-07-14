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
	if node == nil {
		return 0
	}

	pre += node.Val
	l := dfs(node.Left, pre, limit)
	if pre+l < limit {
		node.Left = nil
	}
	r := dfs(node.Right, pre, limit)
	if pre+r < limit {
		node.Right = nil
	}

	max := math.MinInt64
	if node.Left != nil && max < l {
		max = l
	}
	if node.Right != nil && max < r {
		max = r
	}

	if max == math.MinInt64 {
		return node.Val
	}
	return max + node.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
