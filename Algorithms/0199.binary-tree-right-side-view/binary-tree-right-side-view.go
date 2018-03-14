package problem0199

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	// 递归求解
	ls := rightSideView(root.Left)
	rs := rightSideView(root.Right)

	// 左边的分支比右边长的部分，还是可以从右边看到
	if len(ls) > len(rs) {
		rs = append(rs, ls[len(rs):]...)
	}

	// 添加 root 节点的值
	res := make([]int, len(rs)+1)
	res[0] = root.Val
	copy(res[1:], rs)

	return res
}
