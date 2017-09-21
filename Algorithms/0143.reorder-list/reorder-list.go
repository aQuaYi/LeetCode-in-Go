package Problem0143

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type ListNode = kit.ListNode

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	next := head.Next
	temp := next
	for temp.Next.Next != nil {
		temp = temp.Next
	}
	end := temp.Next
	temp.Next = nil

	head.Next = end
	end.Next = next

	reorderList(next)
}
