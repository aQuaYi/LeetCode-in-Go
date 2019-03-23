package problem0958

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined.
type TreeNode = kit.TreeNode

func isCompleteTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 1, 128)

	queue[0] = root
	hasSeenNil := false

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				hasSeenNil = true
				continue
			}
			if hasSeenNil {
				return false
			}
			queue = append(queue, node.Left, node.Right)
		}
		queue = queue[size:]
	}

	return true
}
