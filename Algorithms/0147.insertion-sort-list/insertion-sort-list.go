package problem0147

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type ListNode = kit.ListNode

func insertionSortList(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}

	cur := head
	for cur != nil && cur.Next != nil {
		p := cur.Next
		if cur.Val <= p.Val {
			// p 不是需要插入到 cur 之前的元素
			cur = p
			continue
		}

		// p 是需要插入的元素
		// 把 p 从 cur 后删除
		cur.Next = p.Next
		// 要把 p 插入到合适的 pre 和 next 之间
		pre, next := headPre, headPre.Next
		// 合适的位置的意思是
		// pre.Val < p.Val <= next.Val
		for next.Val < p.Val {
			pre = next
			next = next.Next
		}
		// 插入
		pre.Next = p
		p.Next = next
	}

	return headPre.Next
}
