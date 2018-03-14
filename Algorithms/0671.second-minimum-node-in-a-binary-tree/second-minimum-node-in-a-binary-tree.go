package problem0671

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

const intMax = 1<<63 - 1

func findSecondMinimumValue(root *TreeNode) int {
	res := intMax

	helper(root, root.Val, &res)

	if res == intMax {
		return -1
	}
	return res
}

func helper(root *TreeNode, lo int, hi *int) {
	if root == nil {
		return
	}

	if lo < root.Val && root.Val < *hi {
		*hi = root.Val
	}

	helper(root.Left, lo, hi)
	helper(root.Right, lo, hi)
}
