package problem0606

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"strconv"
)

type TreeNode = kit.TreeNode

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	res := strconv.Itoa(t.Val)

	if t.Left == nil && t.Right == nil {
		return res
	}

	res += "("+ tree2str(t.Left)+")"

	if t.Right != nil {
	res += "("+ tree2str(t.Right)+")"	
	}

	return res
}
