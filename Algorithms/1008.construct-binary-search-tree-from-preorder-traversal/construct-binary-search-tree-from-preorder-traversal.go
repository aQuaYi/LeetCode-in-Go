package problem1008

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

// ref: https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/discuss/252232/JavaC%2B%2BPython-O(N)-Solution
func bstFromPreorder(A []int) *TreeNode {
	i, n := 0, len(A)

	var helper func(int) *TreeNode
	helper = func(bound int) *TreeNode {
		if i == n || A[i] > bound {
			return nil
		}
		root := &TreeNode{
			Val: A[i],
		}
		i++
		root.Left = helper(root.Val)
		root.Right = helper(bound)
		return root
	}

	return helper(1 << 32)
}
