package problem0889

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is definited for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

func constructFromPrePost(pre []int, post []int) *TreeNode {
	size := len(pre)

	if size == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if size == 1 {
		return res
	}

	leftVal := pre[1]
	// 查找 leftVal 在 post 中的索引号
	i := 0
	for ; i < size; i++ {
		if post[i] == leftVal {
			break
		}
	}

	res.Left = constructFromPrePost(pre[1:i+2], post[0:i+1])
	res.Right = constructFromPrePost(pre[i+2:], post[i+1:size-1])

	return res
}
