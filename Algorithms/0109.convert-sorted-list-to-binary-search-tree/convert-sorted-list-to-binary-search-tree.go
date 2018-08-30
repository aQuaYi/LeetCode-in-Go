package problem0109

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode is 题目预先定义的数据结构
type ListNode = kit.ListNode

// TreeNode is 题目预先定义的数据结构
type TreeNode = kit.TreeNode

func sortedListToBST(head *ListNode) *TreeNode {
	return sortedChild(head, nil)
}

func sortedChild(begin, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}

	if begin.Next == end {
		return &TreeNode{Val: begin.Val}
	}

	slow, fast := begin, begin
	for fast != end && fast.Next != end {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// at now, slow is the mid of [begin, end)

	return &TreeNode{
		Val:   slow.Val,
		Left:  sortedChild(begin, slow),
		Right: sortedChild(slow.Next, end),
	}
}
