package Problem0144

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func preorderTraversal(root *TreeNode) []int {
    var right []*TreeNode
    var out []int
    for cur := root; cur != nil; {
        out = append(out, cur.Val)
        if cur.Left == nil && cur.Right == nil {
            if len(right) == 0 {
                break
            }
            cur = right[len(right) - 1]
            right = right[:len(right) - 1]
            continue
		}
		
        if cur.Left != nil && cur.Right != nil {
            right = append(right, cur.Right)
		}
		
        if cur.Left != nil {
            cur = cur.Left
            continue
        }
        cur = cur.Right
    }
    return out
}
