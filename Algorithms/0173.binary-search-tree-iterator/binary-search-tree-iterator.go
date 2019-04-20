package problem0173

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
/* Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = kit.TreeNode

// BSTIterator is the iterator of BST
type BSTIterator struct {
	stack []*TreeNode
}

// Constructor returns a BST iterator
func Constructor(root *TreeNode) BSTIterator {
	stack := make([]*TreeNode, 0, 128)
	res := BSTIterator{
		stack: stack,
	}
	res.push(root)
	return res
}

// Next returns the next smallest number
func (it *BSTIterator) Next() int {
	size := len(it.stack)
	var top *TreeNode
	it.stack, top = it.stack[:size-1], it.stack[size-1]
	it.push(top.Right)
	return top.Val
}

// HasNext returns whether we have a next smallest number
func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}

func (it *BSTIterator) push(root *TreeNode) {
	for root != nil {
		it.stack = append(it.stack, root)
		root = root.Left
	}
}
