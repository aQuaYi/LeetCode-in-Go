package problem0142

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

func detectCycle(head *ListNode) *ListNode {
	headPre := &ListNode{
		Next: head,
	}
	return detect(headPre)
}

func detect(head *ListNode) *ListNode {
	hasSeen := make(map[*ListNode]bool)
	for head != nil && head.Next != nil {
		if hasSeen[head.Next] {
			return head.Next
		}
		hasSeen[head.Next] = true
		head = head.Next
	}
	return nil
}
