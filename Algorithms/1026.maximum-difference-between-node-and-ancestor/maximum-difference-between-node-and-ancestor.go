package problem1026

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func maxAncestorDiff(root *TreeNode) int {
	res := 0
	helper(root, root.Val, &res)
	return res
}

func helper(node *TreeNode, val int, res *int) (int, int) {
	if node == nil {
		return val, val
	}
	val = node.Val

	lMin, lMax := helper(node.Left, val, res)
	rMin, rMax := helper(node.Right, val, res)
	cMin, cMax := min(lMin, rMin), max(lMax, rMax)

	*res = max(*res,
		max(abs(val-cMin), abs(val-cMax)),
	)

	return min(val, cMin), max(val, cMax)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
