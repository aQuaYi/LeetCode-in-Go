package problem0538

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func convertBST(root *TreeNode) *TreeNode {
	sum :=0
	travel(root, &sum)
	return root
}

// 从大到小遍历 BST 并沿路更新 sum 和 节点值
func travel(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	travel(root.Right, sum)
	*sum += root.Val
	root.Val= *sum
	travel(root.Left, sum)
}