package problem0894

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode 是题目预定义的树结构
type TreeNode = kit.TreeNode

var forest = [10][]*TreeNode{[]*TreeNode{&TreeNode{Val: 0}}}

func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return nil
	} else if forest[N/2] != nil {
		return forest[N/2]
	}

	trees := make([]*TreeNode, 0, N*2)
	for les := 1; les <= N-2; les += 2 {
		ris := N - 1 - les
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

	forest[N/2] = trees

	return forest[N/2]
}
