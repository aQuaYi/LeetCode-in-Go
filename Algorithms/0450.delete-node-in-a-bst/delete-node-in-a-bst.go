package problem0450

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return root  
	}

    if root.Val == key {
        if root.Left == nil && root.Right == nil {
            return nil
        } else if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        } else {
            max := findMax(root.Left)
            root.Val = max
            root.Left = deleteNode(root.Left, max)
        }
    } else if root.Val > key {
        root.Left = deleteNode(root.Left, key)
    } else {
        root.Right = deleteNode(root.Right, key)
    }
    
    return root
}

func findMax(node *TreeNode)int  {
    if node.Right == nil {
        return node.Val
    }
    return findMax(node.Right)
}

