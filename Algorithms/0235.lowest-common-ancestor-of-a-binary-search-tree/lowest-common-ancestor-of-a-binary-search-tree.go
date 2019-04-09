package problem0235

import "github.com/aQuaYi/LeetCode-in-Go/kit"

type TreeNode = kit.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
