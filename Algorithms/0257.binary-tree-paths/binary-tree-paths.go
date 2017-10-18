package Problem0257

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"strconv"
)

type TreeNode = kit.TreeNode

func binaryTreePaths(root *TreeNode) []string {
if root == nil {
	return nil  
}

	res := make([]string, 0, 16)

	var search func(string, *TreeNode)
	search = func(pre string, root *TreeNode) {
		if pre == "" {
			pre = strconv.Itoa(root.Val)
		} else {
			pre = pre + "->" + strconv.Itoa(root.Val)
		}

		if root.Left != nil {
			search(pre, root.Left)
		}

		if root.Right != nil {
			search(pre, root.Right)
		}

		if root.Left == nil && root.Right == nil {
			res = append(res, pre)
		}
	}

	search("", root)

	return res
}
