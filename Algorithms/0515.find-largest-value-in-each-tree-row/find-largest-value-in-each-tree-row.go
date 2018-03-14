package problem0515

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"math"
)

type TreeNode = kit.TreeNode

func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	// 进行 BFS
	queue := make([]*TreeNode, 1, 1024)
	queue[0] = root
	n := 1 // 同一个 row 剩余的 node 个数
	max := math.MinInt64
	for len(queue) > 0 {
		// 取出当前 row 的 node
		node := queue[0]
		queue = queue[1:]
		n--
		// 更新当前行的 max
		if max < node.Val {
			max = node.Val
		}
		// 把下一个 row 的节点，添加入 queue
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		// 当同一个 row 的 node 检查完毕后，
		// 保存数据，并开启新的一轮检查
		if n == 0 {
			n = len(queue)
			res = append(res, max)
			max = math.MinInt64
		}
	}

	return res
}
