package problem0099

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

// 已经假设
//    BST 中没有重复的值
//    root != nil
func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	var dfs func(*TreeNode)
	// dfs 是中序遍历
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}

		if prev != nil && prev.Val > root.Val {
			// 如果要调整 [1, 2, 6, 4, 5, 3, 7] 中错位的 6 和 3
			// 其实是把 [6, 4] 中的较大值与 [5, 3] 中的较小值交换。这时，有两组错序。
			// 但是，还有可能是
			// [1, 3, 2] 中错位的 3 和 2，只有一组错序的 [3, 2]
			// 以下的两个 if 语句，可以兼容以上两种情况
			if first == nil {
				first = prev
			}
			if first != nil {
				// 当存在第二组错序的时候
				// second 的值，会被修改
				second = root
			}
		}

		prev = root

		if root.Right != nil {
			dfs(root.Right)
		}
	}

	dfs(root)

	// 题目已经保证存在被交换的节点了
	// 无需检查 first 和 second 是否为 nil
	first.Val, second.Val = second.Val, first.Val
}
