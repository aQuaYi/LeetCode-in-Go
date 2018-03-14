package problem0637

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0, 128)

	nodes := make([]*TreeNode, 1, 1024)
	nodes[0] = root

	for len(nodes) > 0 {
		n := len(nodes)

		sum := 0
		for i := 0; i < n; i++ {
			sum += nodes[i].Val
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}

		res = append(res,float64( sum)/float64(n))
		nodes = nodes[n:]
	}

	return res
}
