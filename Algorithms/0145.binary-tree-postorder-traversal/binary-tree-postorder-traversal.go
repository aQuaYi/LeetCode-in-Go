package problem0145

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func postorderTraversal(root *TreeNode) []int {
	// 保存以后才处理的节点
	var stack []*TreeNode
	res := []int{}

	cur := root
	for cur != nil {
		if cur.Left == nil && cur.Right == nil { // 到达 leaf 节点
			res = append(res, cur.Val)
			// 空 stack 说明，已经完成了遍历
			if len(stack) == 0 {
				break
			}
			// 从 stack 中取出下个节点
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			// cur 不是 leaf，就先让 cur 入 stack
			// 暂存 cur 的左右节点
			left, right  := cur.Left, cur.Right
			// 裁剪 cur，避免重复访问
			cur.Left, cur.Right = nil, nil
			stack = append(stack, cur)

			// 移动 cur
			if left != nil {
				cur = left
				if right != nil {
					// 非空的 right 入 stack
					stack = append(stack, right)
				}
			} else { // left == nil 
				cur = right
			}
		}
	}

	return res
}
