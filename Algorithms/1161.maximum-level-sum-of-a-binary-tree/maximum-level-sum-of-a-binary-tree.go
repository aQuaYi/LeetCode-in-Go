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
	res, maxSum := 0, math.MinInt64
	queue := make([]*TreeNode, 1, 1<<12)
	queue[0] = root

	// BFS
	level := 0
	for len(queue) > 0 {
		size := len(queue)
		level++
		sum := 0
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

		if maxSum < sum {
			maxSum = sum
			res = level
		}

		queue = queue[size:]
	}

	return res
}
