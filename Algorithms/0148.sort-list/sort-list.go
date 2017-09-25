package Problem0148

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type ListNode = kit.ListNode

func sortList(head *ListNode) *ListNode{
    if head == nil || head.Next == nil {
		return head
	}
	secHead := Split(head)
	return MergeList(sortList(head), sortList(secHead))
}

// 从中间位置，切分 list
func Split(head *ListNode) *ListNode {
    if head == nil||head.Next == nil{
        return nil
	}
	// fast 的变化速度是 slow 的两倍
	// 当 fast 到达末尾的时候，slow 正好在 list 的中间
	slow, fast := head, head
	var tail *ListNode
	for fast != nil && fast.Next != nil {
		tail = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	tail.Next = nil
	return slow
}

func MergeList(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	} else if right == nil {
		return left
	}
	var head, cur, p *ListNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur = left
			left = left.Next
		} else {
			cur = right
			right = right.Next
		}
		if head == nil {
			head = cur
		} else {
			p.Next = cur
		}
		p = cur
	}
	if left == nil {
		p.Next = right
	} else {
		p.Next = left
	}
	return head
}
