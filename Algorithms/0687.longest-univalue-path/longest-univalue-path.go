package Problem0687

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//
func helper(n *TreeNode, maxLen *int) int {
	if n == nil {
		return 1
	}

	l := helper(n.Left, maxLen)
	r := helper(n.Right, maxLen)
	lt := 1

	if n.Left != nil && n.Val == n.Left.Val {
		*maxLen = max(*maxLen, l+1)
		lt = max(lt, l+1)
	}
	if n.Right != nil && n.Val == n.Right.Val {
		*maxLen = max(*maxLen, r+1)
		lt = max(lt, r+1)
	}
	if n.Left != nil && n.Right != nil && n.Val == n.Left.Val && n.Val == n.Right.Val {
		*maxLen = max(*maxLen, l+r+1)
	}

	return lt
}

func longestUnivaluePath(root *TreeNode) int {
	maxLen := 1
	helper(root, &maxLen)
	return maxLen - 1
}
