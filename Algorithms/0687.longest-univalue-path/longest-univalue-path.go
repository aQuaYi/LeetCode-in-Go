package problem0687

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	helper(root, &maxLen)
	return maxLen
}

// 返回从 root 出发拥有相同 Val 值的线路上的 edge 数量
func helper(n *TreeNode, maxLen *int) int {
	if n == nil {
		return 0
	}

	l := helper(n.Left, maxLen)
	r := helper(n.Right, maxLen)
	res := 0

	// 左侧单边的最长距离
	if n.Left != nil && n.Val == n.Left.Val {
		*maxLen = max(*maxLen, l+1)
		res = max(res, l+1)
	}
	// 右侧单边的最长距离
	if n.Right != nil && n.Val == n.Right.Val {
		*maxLen = max(*maxLen, r+1)
		res = max(res, r+1)
	}
	// 通过根节点的最长边
	if n.Left != nil && n.Val == n.Left.Val &&
		n.Right != nil && n.Val == n.Right.Val {
		*maxLen = max(*maxLen, l+r+2)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
