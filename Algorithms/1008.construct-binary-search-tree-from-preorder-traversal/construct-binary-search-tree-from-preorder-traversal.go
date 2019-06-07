package problem1008

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	v, less, more := split(preorder)
	return &TreeNode{
		Val:   v,
		Left:  bstFromPreorder(less),
		Right: bstFromPreorder(more),
	}
}

func split(A []int) (int, []int, []int) {
	h := A[0]
	i := 1
	for i < len(A) && A[i] < h {
		i++
	}
	return h, A[1:i], A[i:]
}
