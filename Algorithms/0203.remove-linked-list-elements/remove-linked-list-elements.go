package Problem0203

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type ListNode = kit.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	headPre := &ListNode{Next: head}

	temp := headPre
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return headPre.Next
}
