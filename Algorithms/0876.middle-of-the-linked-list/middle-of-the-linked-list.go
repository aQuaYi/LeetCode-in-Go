package problem0876

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode defined for singly-linked list.
// type ListNode struct {
//     Val int
//     Next *ListNode
// }
type ListNode = kit.ListNode

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
