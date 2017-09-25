package Problem0147

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type ListNode = kit.ListNode

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	count := 0
	temp := &ListNode{Next: head}
	for temp.Next != nil {
		count++
		temp = temp.Next
	}

	headPre := &ListNode{Next: head}

	for cnt := count - 1; cnt > 0; cnt-- {
		temp = headPre
		for c := cnt; c > 0; c-- {
			if temp.Next.Val > temp.Next.Next.Val {
				tempNext := temp.Next
				temp.Next = temp.Next.Next
				tempNext.Next = temp.Next.Next
				temp.Next.Next = tempNext
			}
			temp = temp.Next
		}
	}

	return headPre.Next
}
