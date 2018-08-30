package problem0109

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode is 题目预先定义的数据结构
type ListNode = kit.ListNode

// TreeNode is 题目预先定义的数据结构
type TreeNode = kit.TreeNode

func sortedListToBST(head *ListNode) *TreeNode {
	return transMidToRoot(head, nil)
}

func transMidToRoot(begin, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}

	if begin.Next == end {
		return &TreeNode{Val: begin.Val}
	}

	fast, slow := begin, begin
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow

	return &TreeNode{
		Val:   mid.Val,
		Left:  transMidToRoot(begin, mid),
		Right: transMidToRoot(mid.Next, end),
	}
}
