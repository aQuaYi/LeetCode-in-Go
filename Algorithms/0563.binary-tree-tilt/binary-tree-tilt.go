package Problem0563

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func findTilt(root *TreeNode) int {
	if root ==nil {
		return 0
	}

	l:= sum(root.Left)
	r:= sum(root.Right)

	if l> r {
		return l-r
	}
	return r-l
}

func sum (root *TreeNode)int {
if root == nil {
	return 0
}	

return root.Val+ sum(root.Left)+ sum(root.Right)
}