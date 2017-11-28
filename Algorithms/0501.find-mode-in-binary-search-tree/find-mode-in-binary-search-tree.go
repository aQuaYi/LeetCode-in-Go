package Problem0501

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode
// TODO: 这个方法，完全没有利用上 BST 的特性
// 改进这个方法，效率提升一倍
func findMode(root *TreeNode) []int {
	r := map[int]int{}
	search(root, r)

	max := -1
	for _, v := range r {
		if max < v {
			max = v
		}
	}

	res := []int{}
	for n, v := range r {
		if v == max {
			res = append(res, n)
		}
	}

	return res
}

func search(root *TreeNode, rec map[int]int) {
	if root == nil {
		return
	}

	(rec)[root.Val]++

	search(root.Left, rec)
	search(root.Right, rec)
}
