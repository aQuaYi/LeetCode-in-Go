package Problem0515

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"math"
)

type TreeNode = kit.TreeNode

func largestValues(root *TreeNode) []int {
	res := []int{}
if root == nil {
	return res
}

	queue := []*TreeNode{root}
	n := 1
	max := math.MinInt64

	for len(queue)>0 {
		node := queue[0]
		queue = queue[1:]
		n--

		if max < node.Val {
			max = node.Val
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		if n == 0 {
			n = len(queue)
			res = append(res, max)
			max = math.MinInt64
		}
	}

	return res
}
