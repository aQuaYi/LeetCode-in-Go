package Problem0652

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"strconv"
)

type TreeNode = kit.TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	roots := make(map[string]int, 1024)
	res := make([]*TreeNode, 0, 1024)

	helper(root, roots, &res)

	return res
}

func helper(root *TreeNode, roots map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return ""
	}

	l := helper(root.Left, roots, res)
	r := helper(root.Right, roots, res)

	key := strconv.Itoa(root.Val) + "(" + l + ")(" + r + ")"

	roots[key]++
	if roots[key] == 2 {
		*res = append(*res, root)
	}

	return key
}
