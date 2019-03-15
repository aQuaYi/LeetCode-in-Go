package problem0951

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined in kit library
type TreeNode = kit.TreeNode

func flipEquiv(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if (r1 != nil && r2 == nil) ||
		(r1 == nil && r2 != nil) ||
		r1.Val != r2.Val {
		return false
	}

	if (flipEquiv(r1.Left, r2.Left) && flipEquiv(r1.Right, r2.Right)) ||
		(flipEquiv(r1.Left, r2.Right) && flipEquiv(r1.Right, r2.Left)) {
		return true
	}

	return false
}
