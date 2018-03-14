package problem0144

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func preorderTraversal(root *TreeNode) []int {
	// rightStack 用来暂存右侧节点
	var rightStack []*TreeNode
	var res []int

	for cur := root; cur != nil; {
		res = append(res, cur.Val)

		if cur.Left != nil {
			if cur.Right != nil {
				rightStack = append(rightStack, cur.Right)
			}
			cur = cur.Left
		} else { // cur.Left == nil 
			if cur.Right != nil {
				cur = cur.Right
			} else { // cur.Left == cur.Right == nil
				// stack 已空
				// 说明已经完成遍历了
				if len(rightStack) == 0 {
					break
                }
                // 否则
				// 取出最后放入的右侧节点，继续 for 循环
				cur = rightStack[len(rightStack)-1]
				rightStack = rightStack[:len(rightStack)-1]
			}
		}
    }

	return res
}
