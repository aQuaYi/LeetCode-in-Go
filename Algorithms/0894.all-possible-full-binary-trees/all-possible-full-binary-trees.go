package problem0894

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode 是题目预定义的树结构
type TreeNode = kit.TreeNode

var forest = make([][]*TreeNode, 20)

func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return nil
	}

	if forest[N] != nil {
		return forest[N]
	}

	forest[1] = []*TreeNode{&TreeNode{Val: 0}}

	for n := 3; n <= N; n += 2 {
		trees := make([]*TreeNode, 0, n*2)
		for les := 1; les <= n-2; les += 2 {
			ris := n - 1 - les
			for _, left := range allPossibleFBT(les) {
				for _, right := range allPossibleFBT(ris) {
					trees = append(trees, &TreeNode{
						Val:   0,
						Left:  left,
						Right: right,
					})
				}
			}
		}
		forest[n] = trees
	}

	return forest[N]
}
