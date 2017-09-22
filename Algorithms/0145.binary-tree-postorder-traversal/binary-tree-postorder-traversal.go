package Problem0145

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	res := []int{}

	cur := root
	for cur != nil {
		if cur.Left == nil && cur.Right == nil { // 到达 leaf 节点
			res = append(res, cur.Val)

			if len(stack) == 0 {
				break
			}
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			left := cur.Left
			right := cur.Right
			cur.Left, cur.Right = nil, nil
			stack = append(stack, cur)

			if left != nil {
				cur = left
				if right != nil {
					stack = append(stack, right)
				}
			} else {
				cur = right
			}
		}
	}
	return res
}
