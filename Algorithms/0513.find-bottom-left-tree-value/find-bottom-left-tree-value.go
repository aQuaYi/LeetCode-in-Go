package Problem0513

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	rec := make([]int, 2)
	helper(root, 1,  rec)
	return rec[1]
}

func helper(root *TreeNode, row int,   rec []int) {
	if root == nil {
		return
	}

	if rec[0] < row {
		rec[0] = row
		rec[1] = root.Val
	}

	helper(root.Left, row+1,  rec)
	helper(root.Right, row+1,  rec)
}
