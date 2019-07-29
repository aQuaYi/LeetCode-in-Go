package problem1110

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined
type TreeNode = kit.TreeNode

func delNodes(root *TreeNode, delete []int) []*TreeNode {
	newRoot := &TreeNode{
		Val:  0,
		Left: root,
	}

	isDeleted := [10001]bool{}
	isDeleted[0] = true
	for _, d := range delete {
		isDeleted[d] = true
	}

	res := make([]*TreeNode, 0, 1000)

	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		//
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		//
		if !isDeleted[node.Val] {
			return node
		}
		//
		if node.Left != nil {
			res = append(res, node.Left)
		}
		//
		if node.Right != nil {
			res = append(res, node.Right)
		}
		//
		return nil
	}

	dfs(newRoot)

	return res
}
