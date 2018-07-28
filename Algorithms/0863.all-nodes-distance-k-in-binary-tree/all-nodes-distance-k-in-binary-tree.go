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

// walk 返回从 root 到 target 的距离
// 返回值 <0 表示 target 不在 root 下方
func walk(root, target *TreeNode, res *[]int, dist, k int) int {
	if root == nil {
		return -1
	}

	if root == target {
		walk(root.Left, nil, res, 1, k)
		walk(root.Right, nil, res, 1, k)
		return 0
	}

	// dist < 0 表示还没有找到 target 节点
	if dist < 0 {
		// 先检查左边
		if l := walk(root.Left, target, res, -1, k); l >= 0 {
			// 此时 target 在 root.Left 下方
			if dist = l + 1; dist == k {
				// root 到 target 的距离为 k
				// 所以要把 root.Val 放入 res
				*res = append(*res, root.Val)
			} else if dist < k {
				// root 到 target 的距离不到 k
				// 但是 root.Right 到 target 的距离为 dist+1
				// 所以，从 root.Right 以 dist+1 的距离继续向下寻找
				walk(root.Right, nil, res, dist+1, k)
			}
		} else if r := walk(root.Right, target, res, -1, k); r >= 0 {
			// 右边与左边对称
			if dist = r + 1; dist == k {
				*res = append(*res, root.Val)
			} else if dist < k {
				walk(root.Left, nil, res, dist+1, k)
			}
		}
		return dist
	}

	if dist == k {
		*res = append(*res, root.Val)
	} else if dist < k {
		walk(root.Left, nil, res, dist+1, k)
		walk(root.Right, nil, res, dist+1, k)
	}

	return dist
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}

	res := make([]int, 0, 2048)
	walk(root, target, &res, -1, k)

	return res
}
