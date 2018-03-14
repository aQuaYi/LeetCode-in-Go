package problem0669

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)
type TreeNode = kit.TreeNode
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil 
	}

	if root.Val < L {
return trimBST(root.Right, L, R)
	}

	if R< root.Val  {
return trimBST(root.Left, L, R)
	}

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
