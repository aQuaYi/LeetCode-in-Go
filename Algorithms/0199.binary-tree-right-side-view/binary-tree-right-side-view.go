package Problem0199

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil 
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	ls := rightSideView(root.Left)
	rs := rightSideView(root.Right)

	if len(ls ) > len(rs) {
		rs = append(rs, ls[len(rs):]...)
	}

	res := make([]int , len(rs)+1)
	res[0] = root.Val
	copy(res[1:], rs)

	return res 
}
