package problem0160

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hasSeen := make(map[*ListNode]bool, 128)
	for headA != nil {
		hasSeen[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if hasSeen[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
