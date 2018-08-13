package problem0872

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode definition for a binary tree node.
//  type TreeNode struct {
//      Val int
//      Left *TreeNode
//      Right *TreeNode
//  }
type TreeNode = kit.TreeNode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// 因为题目说了，root 的全部节点不会超过 100 个
	// 那么 leaf 节点就更不会超过 100 个了
	// 所以 a 始终是 s 的底层数组，不会发生变更
	// 结尾比较 a1 和 a2 的异同就会很方便

	a1 := [100]int{}
	s1 := a1[:0]
	search(root1, &s1)

	a2 := [100]int{}
	s2 := a2[:0]
	search(root2, &s2)

	return a1 == a2
}

func search(root *TreeNode, sp *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*sp = append(*sp, root.Val)
		return
	}
	search(root.Left, sp)
	search(root.Right, sp)
}
