package problem0652

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"strconv"
)

type TreeNode = kit.TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	count := make(map[string]int, 1024)
	res := make([]*TreeNode, 0, 1024)

	helper(root, count, &res)

	return res
}

func helper(root *TreeNode, count map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "*"
	}

	l := helper(root.Left, count, res)
	r := helper(root.Right, count, res)

	key := strconv.Itoa(root.Val) +  l +  r 

	count[key]++
	if count[key] == 2 {
		*res = append(*res, root)
	}

	return key
}
