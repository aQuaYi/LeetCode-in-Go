package problem1161

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

func maxLevelSum(root *TreeNode) int {
	queue := make([]*TreeNode, 1, 1<<10)
	queue[0] = root

	res, max := 0, math.MinInt64
	level, sum := 1, 0
	// BFS
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if max < sum {
			res, max = level, sum
		}

		level, sum = level+1, 0

		queue = queue[size:]
	}

	return res
}
