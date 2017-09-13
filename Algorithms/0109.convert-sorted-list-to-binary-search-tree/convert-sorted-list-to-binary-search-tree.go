package Problem0109

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type ListNode = kit.ListNode
type TreeNode = kit.TreeNode

func sortedListToBST(head *ListNode) *TreeNode {
	return sortedArrayToBST(list2Slice(head))
}

func list2Slice(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}
