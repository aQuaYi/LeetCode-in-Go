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
	if rescur(root, voyage, &res) {
		return res
	}
	return []int{-1}
}

func rescur(root *TreeNode, voyage []int, res *[]int) bool {
	if root == nil {
		if len(voyage) == 0 {
			return true
		}
		return false
	}

	if len(voyage) == 1 && root.Val == voyage[0] && root.Left == nil && root.Right == nil {
		return true
	}

	if len(voyage) == 0 || root.Val != voyage[0] {
		return false
	}

	if len(voyage) < 2 && root.Left != nil {
		return false
	}

	if voyage[1] != root.Left.Val {
		*res = append(*res, root.Val)
		root.Left, root.Right = root.Right, root.Left
	}

	right := 0
	if root.Right != nil {
		right = root.Right.Val
	}

	l, r := split(voyage, right)

	return rescur(root.Left, l, res) && rescur(root.Right, r, res)
}

func split(voyage []int, right int) ([]int, []int) {
	begin := 1
	for ; begin < len(voyage); begin++ {
		if voyage[begin] == right {
			break
		}
	}
	return voyage[1:begin], voyage[begin:]
}
