package problem0863

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*TreeNode 是题目预定义的类型
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = kit.TreeNode

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	res := make([]int, 0, 2048)
	search(root, target, k, &res)
	return res
}

// check 检查 root 到 target 的 dist 是否已经是 k 了
func check(root *TreeNode, dist, k int, res *[]int) {
	if root == nil || dist > k {
		return
	}

	if dist == k {
		*res = append(*res, root.Val)
		return
	}

	check(root.Left, dist+1, k, res)
	check(root.Right, dist+1, k, res)
}

// search 返回从 root 到 target 的距离
// 返回值 -1 表示 target 不在 root 及其子节点中
func search(root, target *TreeNode, k int, res *[]int) (dist int) {
	if root == nil {
		return -1
	}

	if root == target {
		check(root, 0, k, res)
		return 0
	}

	isIn := func(child, theOther *TreeNode) bool {
		childDist := search(child, target, k, res)
		if childDist > -1 {
			dist = childDist + 1
			if dist == k {
				*res = append(*res, root.Val)
			} else if dist < k {
				check(theOther, dist+1, k, res)
			}
			return true
		}
		return false
	}

	if !isIn(root.Left, root.Right) && !isIn(root.Right, root.Left) {
		return -1
	}

	return dist
}
