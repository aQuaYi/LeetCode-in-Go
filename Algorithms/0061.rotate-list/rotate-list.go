package Problem0061

import (
	"fmt"
)

// ListNode 是题目预定义的节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	end := head
	cnt := 1
	for end.Next != nil {
		end = end.Next
		cnt++
	}
	fmt.Println(cnt)

	end.Next = head
	fmt.Println(l2s(end))

	cnt -= k % cnt

	for cnt > 1 {
		head = head.Next
		end = end.Next
		cnt--
	}

	end.Next = nil

	return head
}
