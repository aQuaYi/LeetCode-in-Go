package Problem0530

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)
type TreeNode = kit.TreeNode

func getMinimumDifference(root *TreeNode) int {
	res := 1<<31-1 

	if root.Left == nil && root.Right == nil {
		return res
	}

	if root.Left!= nil {
		res = min(res, getMinimumDifference(root.Left))
		res = min(res, root.Val- maxNode(root.Left))
	}

	if root.Right != nil {
		res = min(res , getMinimumDifference(root.Right))
		res = min(res, minNode(root.Right)- root.Val)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minNode(root *TreeNode) int {
if root.Left==nil {
	return root.Val
}	
return minNode(root.Left)
}

func maxNode(root *TreeNode) int {
if root.Right==nil {
	return root.Val
}	
return maxNode(root.Right)
}