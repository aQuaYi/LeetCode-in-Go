package problem0572

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func isSubtree(s *TreeNode, t *TreeNode) bool {    
    if isSame(s, t) { return true }
    
    if s == nil {return false}
    return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
    if s == nil {return t == nil}
    if t == nil {return false}
    
    if s.Val != t.Val {
        return false
    }
    
    return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
