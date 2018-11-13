package problem0148

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type ListNode = kit.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	secHead := split(head)
	return merge(sortList(head), sortList(secHead))
}

// 从中间位置，切分 list
func split(head *ListNode) *ListNode {
	// head.Next != nil
	// 因为， sortList 已经帮忙检查过了

	// fast 的变化速度是 slow 的两倍
	// 当 fast 到达末尾的时候，slow 正好在 list 的中间
	slow, fast := head, head
	var tail *ListNode
	for fast != nil && fast.Next != nil {
		tail = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 斩断 list
	tail.Next = nil
	return slow
}

// 把已经排序好了的两个 list left 和 right
// 进行合并
func merge(left, right *ListNode) *ListNode {
	// left != nil , right != nil
	// 因为， sortList 已经帮忙检查过了

	cur := &ListNode{}
	headPre := cur
	for left != nil && right != nil {
		// cur 指向 left 和 right 中最小的节点
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}

		cur = cur.Next
	}

	// 处理 left 或 right 中，剩下的节点
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}

	return headPre.Next
}
