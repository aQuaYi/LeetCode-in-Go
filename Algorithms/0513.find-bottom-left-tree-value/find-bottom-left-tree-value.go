package problem0513

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]

		// 先把 Right 边的节点放入 queue 才能保证
		// 最后剩下的 root 是 bottom left
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
	}

	return root.Val
}
