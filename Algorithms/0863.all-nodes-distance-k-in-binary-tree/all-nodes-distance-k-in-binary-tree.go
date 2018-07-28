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
	if k == 0 {
		return []int{target.Val}
	}

	res := make([]int, 0, 2048)
	find(root, target, &res, k)

	return res
}

// dist 是 root 与 target 的距离
// 越往下，距离 target 越远
func walk(root *TreeNode, res *[]int, dist, k int) {
	if root == nil {
		return
	}

	if dist == k {
		*res = append(*res, root.Val)
		return
	}

	walk(root.Left, res, dist+1, k)
	walk(root.Right, res, dist+1, k)
}

// find 返回从 root 到 target 的距离
// 返回值 -1 表示 target 不在 root 下方
func find(root, target *TreeNode, res *[]int, k int) int {
	dist := -1
	if root == nil {
		return dist
	}

	// 1. 检查 root ？= target
	if root == target {
		// k = 0 的情况已经被预先处理了
		walk(root.Left, res, 1, k)
		walk(root.Right, res, 1, k)
		return 0
	}

	// 2. 检查 target 是否在 root 左边吗
	if l := find(root.Left, target, res, k); l >= 0 {
		// 此时 target 在 root.Left 下方
		if dist = l + 1; dist == k {
			// root 到 target 的距离为 k
			// 所以要把 root.Val 放入 res
			*res = append(*res, root.Val)
		} else if dist < k {
			// root 到 target 的距离不到 k
			// 但是 root.Right 到 target 的距离为 dist+1
			// 所以，从 root.Right 以 dist+1 的距离继续向下寻找
			walk(root.Right, res, dist+1, k)
		}
	} else if r := find(root.Right, target, res, k); r >= 0 {
		// 3. 检查 target 是否在 root 右边
		// 右边与左边对称
		if dist = r + 1; dist == k {
			*res = append(*res, root.Val)
		} else if dist < k {
			walk(root.Left, res, dist+1, k)
		}
	}

	return dist
}
