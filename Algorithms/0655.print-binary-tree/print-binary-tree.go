package problem0655

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"strconv"
)

type TreeNode = kit.TreeNode

func printTree(root *TreeNode) [][]string {
	level := getLevel(root)
	size := 1<<uint(level) - 1
	res := make([][]string, level)
	for i := range res {
		res[i] = make([]string, size)
	}

	// loc 是 root　节点的值在 res 中的索引号
	loc := 1<<uint(level-1) - 1
	getRes(root, 0, loc, res)

	return res
}

func getRes(root *TreeNode, i, j int, res [][]string) {
	if root == nil {
		return
	}

	res[i][j] = strconv.Itoa(root.Val)

	level := len(res)
	if level-i-2 < 0 {
		return
	}

	diff := 1 << uint(level-i-2)

	getRes(root.Left, i+1, j-diff, res)
	getRes(root.Right, i+1, j+diff, res)
}

// 返回树的层数
func getLevel(root *TreeNode) (res int) {
	res = 1
	queue := make([]*TreeNode, 1, 1024)
	queue[0] = root

	for {
		hasChild := false
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil || queue[i].Right != nil {
				hasChild = true
				break
			}
		}

		if !hasChild {
			break
		}
		res++

		fQueue := queue[:len(queue)]
		queue = queue[len(queue):]

		for i := 0; i < len(fQueue); i++ {
			if fQueue[i].Left != nil {
				queue = append(queue, fQueue[i].Left)
			}
			if fQueue[i].Right != nil {
				queue = append(queue, fQueue[i].Right)
			}
		}
	}

	return res
}
