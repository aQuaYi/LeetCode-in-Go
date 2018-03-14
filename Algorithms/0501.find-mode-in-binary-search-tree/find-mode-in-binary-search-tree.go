package problem0501

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func findMode(root *TreeNode) []int {
	r := map[int]int{}
	search(root, r)

	max := -1
	res := []int{}
	for n, v := range r {
		if max <= v {
			if max < v {
				max = v
				res = res[0:0]
			}
			res = append(res, n)
		}
	}

	return res
}

func search(root *TreeNode, rec map[int]int) {
	if root == nil {
		return
	}

	rec[root.Val]++

	search(root.Left, rec)
	search(root.Right, rec)
}
