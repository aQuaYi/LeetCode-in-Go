package Problem0543

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := helper(root)

	return res
}

func helper(root *TreeNode) (length, diameter int) {
	if root == nil {
		return 0,0 
	}

	leftLen, leftDia := helper(root.Left)
	rightLen, rightDia := helper(root.Right)

	return max(leftLen, rightLen) + 1, max(leftLen+rightLen, max(leftDia, rightDia))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
