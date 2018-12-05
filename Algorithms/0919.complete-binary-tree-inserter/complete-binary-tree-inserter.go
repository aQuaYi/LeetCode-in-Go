package problem0919

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is defined in kit
type TreeNode = kit.TreeNode

// CBTInserter is ...
type CBTInserter struct {
}

// Constructor is ...
func Constructor(root *TreeNode) CBTInserter {

	return CBTInserter{}
}

// Insert is
func (c *CBTInserter) Insert(v int) int {

	return 1
}

// Get_root is ...
func (c *CBTInserter) Get_root() *TreeNode {

	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
}
