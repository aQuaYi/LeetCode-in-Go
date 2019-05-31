package problem0988

import (
	"strings"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func smallestFromLeaf(root *TreeNode) string {
	res := strings.Repeat("z", 27)
	dfs(root, "", &res)
	return res
}

func dfs(node *TreeNode, s string, res *string) {
	s = string('a'+node.Val) + s

	if len(s) > len(*res) {
		return
	}

	if node.Left == nil && node.Right == nil {
		*res = min(*res, s)
		return
	}

	if node.Left != nil {
		dfs(node.Left, s, res)
	}

	if node.Right != nil {
		dfs(node.Right, s, res)
	}

	return
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}
