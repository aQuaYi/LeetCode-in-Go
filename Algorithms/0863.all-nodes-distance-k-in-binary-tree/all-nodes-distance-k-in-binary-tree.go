package problem0863

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

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
// 返回值 -1 表示 target 不在 root 下方
func search(root, target *TreeNode, k int, res *[]int) (dist int, isFound bool) {
	if root == nil {
		isFound = false
		return
	}

	// 1. 检查 root ？= target
	if root == target {
		dist, isFound = 0, true
		check(root, dist, k, res)
		return
	}

	var childDist int

	// 2. 检查 target 是否在 root 左边
	childDist, isFound = search(root.Left, target, k, res)
	if isFound {
		dist = childDist + 1
		if dist == k {
			*res = append(*res, root.Val)
		}
		check(root.Right, dist+1, k, res)
		return
	}

	// 3. 检查 target 是否在 root 右边
	childDist, isFound = search(root.Right, target, k, res)
	if isFound {
		dist = childDist + 1
		if dist == k {
			*res = append(*res, root.Val)
		}
		check(root.Left, dist+1, k, res)
	}
	return
}
