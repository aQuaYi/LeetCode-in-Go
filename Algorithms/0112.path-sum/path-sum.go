package problem0112

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	
	if root.Left == nil && root.Right == nil {
		// 此时 root 才为 leaf
		return sum == 0
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}