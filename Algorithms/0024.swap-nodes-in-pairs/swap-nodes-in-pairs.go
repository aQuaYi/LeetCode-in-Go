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

	// newHead 指向 head.Next 节点
	newHead := head.Next
	// 让 head.Next 指向转换好了 newHead.Next 节点
	head.Next = swapPairs(newHead.Next)
	// 让 newHead.Next 指向 head 节点
	newHead.Next = head
	// newHead 成为新的 head 节点

	return newHead
}
