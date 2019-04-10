package problem0236

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

// NOTICE: 此解法的时间复杂度是 O(n)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil ||
		root == p ||
		root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	// l=nil 意味着， p 和 q 都 **不在** root.Left 中
	// r=nil 同理。
	// 所以，根据题意， l 和 r 不可能同时为 nil
	if l != nil && r != nil {
		// 此时 p 和 q 分别在 root.Left 和 root.Right 中
		return root
	}
	if l == nil {
		// 此时 p 和 q 在 root.Right 中
		return r
	}
	// 此时 p 和 q 在 root.Left 中
	return l
}
