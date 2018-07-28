package problem0863

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode 是题目预定义的类型
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = kit.TreeNode

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	conn := make(map[int][]int, 2048)
	var connect func(*TreeNode, *TreeNode)
	connect = func(parent, child *TreeNode) {
		if parent != nil && child != nil {
			conn[parent.Val] = append(conn[parent.Val], child.Val)
			conn[child.Val] = append(conn[child.Val], parent.Val)
		}
		if child.Left != nil {
			connect(child, child.Left)
		}
		if child.Right != nil {
			connect(child, child.Right)
		}
	}
	// 把 TreeNode 转换成 图
	connect(nil, root)

	bfs := []int{target.Val}
	hasSeen := make(map[int]bool, 2048)
	hasSeen[target.Val] = true

	for K > 0 {
		K--
		qSize := len(bfs)
		for i := 0; i < qSize; i++ {
			for _, x := range conn[bfs[i]] {
				if !hasSeen[x] {
					bfs = append(bfs, x)
					hasSeen[x] = true
				}
			}
		}
		bfs = bfs[qSize:]
	}

	return bfs
}
