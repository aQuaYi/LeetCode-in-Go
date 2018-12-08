package problem0919

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is defined in kit
type TreeNode = kit.TreeNode

// CBTInserter is ...
type CBTInserter struct {
	root *TreeNode
}

// Constructor is ...
func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{root: root}
}

// Insert is
func (c *CBTInserter) Insert(v int) int {
	nds := []*TreeNode{c.root}

	var parent *TreeNode
	added := false

ALL:
	for len(nds) > 0 {
		size := len(nds)
		for i := 0; i < size; i++ {
			parent = nds[i]
			if parent.Left == nil {
				parent.Left = &TreeNode{Val: v}
				added = true
			} else if parent.Right == nil {
				parent.Right = &TreeNode{Val: v}
				added = true
			}

			if added {
				break ALL
			}
			nds = append(nds, parent.Left, parent.Right)
		}
		nds = nds[size:]
	}
	return parent.Val
}

// Get_root is ...
func (c *CBTInserter) Get_root() *TreeNode {
	return c.root
}
