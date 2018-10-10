package problem0897

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	var head = &TreeNode{}
	tail := head
	var rec func(root *TreeNode)
	rec = func(root *TreeNode) {
		if root == nil {
			return
		}
		rec(root.Left)
		root.Left = nil               // 切断 root 与其 Left 的连接，避免形成环
		tail.Right, tail = root, root // 把 root 接上 tail，并保持 tail 指向尾部
		rec(root.Right)
	}
	rec(root)
	return head.Right
}
