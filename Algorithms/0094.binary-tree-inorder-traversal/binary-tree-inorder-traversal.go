package problem0094

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}
