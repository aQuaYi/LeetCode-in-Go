package Problem0098

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func isValidBST(root *TreeNode) bool {
	MIN,MAX := -1<<63,1<<63-1
	
	return cur(MIN, MAX, root) 
}
func cur(min,max int , root *TreeNode) bool  {
	if root == nil {
		return true 
	}
	
	if min < root.Val && root.Val < max  {
		return cur(min, root.Val, root.Left) && cur(root.Val,max, root.Right)
	}
	
	return false  
}