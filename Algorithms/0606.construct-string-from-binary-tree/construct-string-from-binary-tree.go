package Problem0606

import (
	"fmt"
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"strconv"
)

type TreeNode = kit.TreeNode

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}

	if t.Right == nil {
		return fmt.Sprintf("%d(%s)", t.Val, tree2str(t.Left))
	}

	return fmt.Sprintf("%d(%s)(%s)", t.Val, tree2str(t.Left), tree2str(t.Right))
}
