package problem0002

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode defines for singly-linked list.
//  type ListNode struct {
//      Val int
//      Next *ListNode
//  }
type ListNode = kit.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}
