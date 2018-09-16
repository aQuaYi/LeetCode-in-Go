package problem0894

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode 是题目预定义的树结构
type TreeNode = kit.TreeNode

func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return nil
	}

	cache := [20][]*TreeNode{}

	cache[1] = []*TreeNode{&TreeNode{Val: 0}}

	for n := 3; n <= N; n += 2 {
		tmp := make([]*TreeNode, 0, n*2)
		for le := 1; le <= n-2; le += 2 {
			ri := n - 1 - le
			for _, left := range cache[le] {
				for _, right := range cache[ri] {
					tmp = append(tmp, &TreeNode{
						Val:   0,
						Left:  left,
						Right: right,
					})
				}
			}
		}
		cache[n] = tmp
	}

	return cache[N]
}
