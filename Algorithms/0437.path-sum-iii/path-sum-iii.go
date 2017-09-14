package Problem0437

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func pathSum(root *TreeNode, sum int) int {
	res := 0
	if root == nil {
		return res
	}

	newSum := sum - root.Val
	if newSum == 0 {
		res++
		return 1 + pathSum(root.Left, sum) + pathSum(root.Right, sum)
	}

	return pathSum(root.Left, sum) +
		pathSum(root.Right, sum) +
		pathSum(root.Left, newSum) +
		pathSum(root.Right, newSum)
}
