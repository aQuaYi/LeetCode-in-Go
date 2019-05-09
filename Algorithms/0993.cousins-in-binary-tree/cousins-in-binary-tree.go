package problem0993

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is Pre-defined....
type TreeNode = kit.TreeNode

func isCousins(root *TreeNode, x int, y int) bool {
	root = &TreeNode{Left: root}
	px, dx := dfs(root, x)
	py, dy := dfs(root, y)
	return dx == dy && px != py
}

// dfs do NOT check the first root
func dfs(root *TreeNode, x int) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	if (root.Left != nil && root.Left.Val == x) ||
		(root.Right != nil && root.Right.Val == x) {
		return root, 1
	}

	if parent, depth := dfs(root.Left, x); depth > 0 {
		return parent, depth + 1
	}

	if parent, depth := dfs(root.Right, x); depth > 0 {
		return parent, depth + 1
	}

	return nil, 0
}
