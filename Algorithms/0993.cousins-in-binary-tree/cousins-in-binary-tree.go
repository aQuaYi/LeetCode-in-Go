package problem0993

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is Pre-defined....
type TreeNode = kit.TreeNode

func isCousins(root *TreeNode, x int, y int) bool {
	px, dx := find(root, x)
	py, dy := find(root, y)
	return dx == dy && px != py
}

// root's parent is nil
func find(root *TreeNode, x int) (*TreeNode, int) {
	if root == nil {
		return nil, -1
	}
	if root.Val == x {
		return nil, 0
	}

	p, d := find(root.Left, x)
	if d == 0 {
		return root, 1
	} else if d > 0 {
		return p, d + 1
	}

	p, d = find(root.Right, x)
	if d == 0 {
		return root, 1
	} else if d > 0 {
		return p, d + 1
	}

	return nil, -1
}
