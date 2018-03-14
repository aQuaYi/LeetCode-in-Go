package problem0662

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil  {
		return 0
	}

	// locQueue[i] = j 表示
	// 当 树 是满的时候，
	// nodeQueue[i] 节点在其所在层的索引号为 j
	locQueue := make([]int, 1, 1024)
	nodeQueue := make([]*TreeNode, 1, 1024)
	locQueue[0] = 0
	nodeQueue[0] = root

	res := 0

	// 使用 BFS 来遍历树
	for len(locQueue) > 0 {
		res = max(res, getWidth(locQueue))

		n := len(locQueue)

		fLocQueue := locQueue[:n]
		fNodeQueue := nodeQueue[:n]
		locQueue = locQueue[n:]
		nodeQueue = nodeQueue[n:]

		for i := 0; i < n; i++ {
			if fNodeQueue[i].Left != nil {
				nodeQueue = append(nodeQueue, fNodeQueue[i].Left)
				locQueue = append(locQueue, fLocQueue[i]*2)
			}
			if fNodeQueue[i].Right != nil {
				nodeQueue = append(nodeQueue, fNodeQueue[i].Right)
				locQueue = append(locQueue, fLocQueue[i]*2+1)
			}
		}

	}

	return res
}
// 本题也可以使用 DFS 来做

func getWidth(locQueue []int) int {
	n := len(locQueue)
	return locQueue[n-1] - locQueue[0] + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
