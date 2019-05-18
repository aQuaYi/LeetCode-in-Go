package problem0971

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = kit.TreeNode

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	res := make([]int, 0, len(voyage))
	index := 0
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Val != voyage[index] {
			return false
		}
		index++
		if node.Left != nil && node.Left.Val != voyage[index] {
			res = append(res, node.Val)
			return dfs(node.Right) && dfs(node.Left)
		}
		return dfs(node.Left) && dfs(node.Right)
	}
	if dfs(root) {
		return res
	}
	return []int{-1}
}
