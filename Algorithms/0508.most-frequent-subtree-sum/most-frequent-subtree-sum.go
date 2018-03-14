package problem0508

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func findFrequentTreeSum(root *TreeNode) []int {
	rec := make(map[int]int)
	sum(root, rec)

	res := make([]int, 0, len(rec))
	max := -1 
	for s, c := range rec {
		if max <= c {
			if max < c {
				max = c
				res = res[0:0]
			}
			res = append(res , s)
		}
	}

	return res
}

// 递归求子树的和，并统计各个和出现的次数
func sum(root *TreeNode, r map[int]int) int {
	if root == nil {
		return 0
	}

	s := root.Val + sum(root.Left, r) + sum(root.Right, r)
	r[s]++

	return s
}
