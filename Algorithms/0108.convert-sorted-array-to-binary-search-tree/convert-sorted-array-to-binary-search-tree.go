package Problem0108

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	return &TreeNode{
		Val:   nums[n/2],
		Left:  sortedArrayToBST(nums[:n/2]),
		Right: sortedArrayToBST(nums[n/2+1:]),
	}
}
