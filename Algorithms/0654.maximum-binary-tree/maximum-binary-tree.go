package Problem0654

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	v, i := getMax(nums)

	return &TreeNode{
		Val:   v,
		Left:  constructMaximumBinaryTree(nums[:i]),
		Right: constructMaximumBinaryTree(nums[i+1:]),
	}
}

func getMax(nums []int) (val, index int) {
	val = nums[0]
	index = 0
	for i := 1; i < len(nums); i++ {
		if val < nums[i] {
			val = nums[i]
			index = i
		}
	}
	return val, index
}
