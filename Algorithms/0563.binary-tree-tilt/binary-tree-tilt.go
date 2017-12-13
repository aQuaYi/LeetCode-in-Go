package Problem0563

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func findTilt(root *TreeNode) int {
	if root ==nil {
		return 0
	}

var l, r int
helper(root, &l, &r)

	if l> r {
		return l-r
	}
	return r-l
}

func helper (root *TreeNode, l, r *int) {
if root == nil {
	return 
}	

helper(root.Left, l, r)
helper(root.Right, l, r)

if root.Left !=nil {
	*l += root.Left.Val
}

if root.Right != nil {
	*r+= root.Right.Val
}
}