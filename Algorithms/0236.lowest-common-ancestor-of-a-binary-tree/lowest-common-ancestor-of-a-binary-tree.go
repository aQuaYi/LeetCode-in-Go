package problem0236

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	isAncestorOfP := makeReferee(p)
	if isAncestorOfP(q) {
		return q
	}

	isAncestorOfQ := makeReferee(q)
	if isAncestorOfQ(p) {
		return p
	}

	for {
		if isAncestorOfP(root.Left) &&
			isAncestorOfQ(root.Left) {
			root = root.Left
			continue
		}
		if isAncestorOfP(root.Right) &&
			isAncestorOfQ(root.Right) {
			root = root.Right
			continue
		}
		break
	}

	return root
}

func makeReferee(child *TreeNode) func(*TreeNode) bool {
	rec := make(map[int]bool, 1024)
	var f func(*TreeNode) bool
	f = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		res, ok := rec[root.Val]
		if ok {
			return res
		}
		res = (root == child) || f(root.Left) || f(root.Right)
		rec[root.Val] = res
		return res
	}
	return f
}
