package problem0173

import "github.com/aQuaYi/LeetCode-in-Go/kit"

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
}

// Constructor returns a BST iterator
func Constructor(root *TreeNode) BSTIterator {

	return BSTIterator{}
}

// Next returns the next smallest number */
func (it *BSTIterator) Next() int {

	return 0
}

// HasNext returns whether we have a next smallest number */
func (it *BSTIterator) HasNext() bool {

	return false
}
