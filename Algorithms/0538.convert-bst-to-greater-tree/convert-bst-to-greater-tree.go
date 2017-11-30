package Problem0538

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func convertBST(root *TreeNode) *TreeNode {
	sum :=0
	travel(root, &sum)
	return root
}

func travel(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	travel(root.Right, sum)
	
	root.Val+= *sum

	*sum = root.Val

	travel(root.Left, sum)
}