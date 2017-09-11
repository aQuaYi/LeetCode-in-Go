package kit

import (
	"fmt"
)

func preIn2Tree(preorder, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	if len(inorder) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: preorder[0],
	}

	if len(inorder) == 1 {
		return res
	}

	idx := indexOf(res.Val, inorder)

	res.Left = preIn2Tree(preorder[1:idx+1], inorder[:idx])
	res.Right = preIn2Tree(preorder[idx+1:], inorder[idx+1:])

	return res
}

func inPost2Tree(inorder, postorder []int) *TreeNode {
	if len(postorder) != len(inorder) {
		panic("inPost2Tree 中两个切片的长度不相等")
	}

	if len(inorder) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: postorder[len(postorder)-1],
	}

	if len(inorder) == 1 {
		return res
	}

	idx := indexOf(res.Val, inorder)

	res.Left = inPost2Tree(inorder[:idx], postorder[:idx])
	res.Right = inPost2Tree(inorder[idx+1:], postorder[idx:len(postorder)-1])

	return res
}

func tree2preorder(root *TreeNode) []int {
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, tree2preorder(root.Left)...)
	res = append(res, tree2preorder(root.Right)...)

	return res
}

func tree2inorder(root *TreeNode) []int {
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := tree2inorder(root.Left)
	res = append(res, root.Val)
	res = append(res, tree2inorder(root.Right)...)

	return res
}

func tree2postorder(root *TreeNode) []int {
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := tree2postorder(root.Left)
	res = append(res, tree2postorder(root.Right)...)
	res = append(res, root.Val)

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}
