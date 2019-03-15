package problem0951

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined in kit library
type TreeNode = kit.TreeNode

func flipEquiv(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if (root1 != nil && root2 == nil) ||
		(root1 == nil && root2 != nil) ||
		root1.Val != root2.Val {
		return false
	}

	if (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)) {
		return true
	}

	return false
}
