package Problem0671

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func findSecondMinimumValue(root *TreeNode) int { 
	if root == nil ||
	root.Left == nil {
		return -1
	}

	if root.Val ==  root.Left.Val && root.Val == root.Right.Val {
		return min(findSecondMinimumValue(root.Left),findSecondMinimumValue(root.Right))
	}

	if root.Left.Val > root.Right.Val {
		return helper(root.Left, root.Right)
	}
		return helper(root.Right, root.Left)
}

func helper(l,s *TreeNode) int {
lv := l.Val
sv := findSecondMinimumValue(s)	
if sv ==-1 {
	return lv
}

return min(lv,sv)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

