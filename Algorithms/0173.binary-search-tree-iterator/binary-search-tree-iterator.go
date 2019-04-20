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
	index *int
	nums  *[]int
}

// Constructor returns a BST iterator
func Constructor(root *TreeNode) BSTIterator {
	nums := convert(root)
	index := 0
	return BSTIterator{
		index: &index,
		nums:  &nums,
	}
}

// Next returns the next smallest number
func (it *BSTIterator) Next() int {
	res := (*it.nums)[*it.index]
	*it.index++
	return res
}

// HasNext returns whether we have a next smallest number
func (it *BSTIterator) HasNext() bool {
	return *it.index < len(*it.nums)
}

func convert(root *TreeNode) []int {
	res := make([]int, 0, 128)
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, res)
	*res = append(*res, root.Val)
	helper(root.Right, res)
}
