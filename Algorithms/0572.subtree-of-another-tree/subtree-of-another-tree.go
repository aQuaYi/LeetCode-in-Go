package Problem0572

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return isSub(s, t, true)
}

func isSub(s, t *TreeNode, isRoot bool) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val && isSub(s.Left, t.Left, false) && isSub(s.Right, t.Right, false) {
		return true
	}

	if !isRoot {
		return false
	}

	return isSub(s.Left, t, true) || isSub(s.Right, t, true)
}
