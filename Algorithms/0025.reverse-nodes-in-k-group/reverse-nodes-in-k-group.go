package problem0025

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode defines for singly-linked list.
type ListNode = kit.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k < 2 {
		return head
	}

	tail, nr := needReverse(head, k)
	if nr {
		nextHead := tail.Next
		head, tail := reverse(head, tail)
		// 递归
		// 把整理好了的前k个节点的尾部，指向整理好了的后面节点的head
		tail.Next = reverseKGroup(nextHead, k)
		return head
	}

	return head
}

func needReverse(head *ListNode, k int) (*ListNode, bool) {
	for k > 1 && head != nil {
		head = head.Next
		k--
	}
	return head, k == 1 && head != nil
}

// // 返回逆转后的首尾节点
// func reverse(head *ListNode) (first, last *ListNode) {
// 	if head == nil || head.Next == nil {
// 		return head, nil
// 	}
// 	gotLast := false
// 	for head != nil {
// 		temp := head.Next
// 		head.Next = first
// 		first = head
// 		head = temp
// 		if !gotLast {
// 			last = first
// 			gotLast = true
// 		}
// 	}
// 	return first, last
// }

// 返回逆转后的首尾节点
func reverse(head, tail *ListNode) (first, last *ListNode) {
	curPre, cur := head, head.Next
	for cur != nil {
		curPre, cur, cur.Next = cur, cur.Next, curPre
	}
	return tail, head
}

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}
