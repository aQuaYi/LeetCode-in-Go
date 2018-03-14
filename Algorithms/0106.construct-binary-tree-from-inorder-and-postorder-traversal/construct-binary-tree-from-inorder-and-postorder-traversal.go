package problem0106

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func buildTree(in []int, post []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: post[len(post)-1],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = buildTree(in[:idx], post[:idx])
	res.Right = buildTree(in[idx+1:], post[idx:len(post)-1])

	return res
}


func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}
