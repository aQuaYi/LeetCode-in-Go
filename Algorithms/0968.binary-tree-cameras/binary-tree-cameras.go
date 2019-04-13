package problem0968

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func minCameraCover(root *TreeNode) int {
	res := 0
	r := helper(root, &res)
	if r == 3 || r == 0 {
		res++
	}
	return res
}

// return
// -1: root is nil
//  0: root is leaf
//  1: root has camera
//  2: root is watched by camera
//  3: root need watched by father's camera
func helper(root *TreeNode, res *int) int {
	if root == nil {
		return -1
	}
	l, r := helper(root.Left, res), helper(root.Right, res)
	if l == -1 && r == -1 {
		return 0
	}
	if l == 0 || r == 0 || l == 3 || r == 3 {
		*res++
		return 1
	}
	if l == 1 || r == 1 {
		return 2
	}
	if l == 2 || r == 2 {
		return 3
	}
	return 4 // TODO: 删除此处内容
}

// 把较深的子树交换到 Left 这边
func trans(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := trans(root.Left), trans(root.Left)
	if l < r {
		root.Left, root.Right = root.Right, root.Left
	}
	d := max(l, r) + 1
	root.Val = d
	return d
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
