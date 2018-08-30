package problem0024

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode is definition for singly-linked list
type ListNode = kit.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}
