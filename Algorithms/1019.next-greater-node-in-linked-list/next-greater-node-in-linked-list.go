package problem1019

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

func nextLargerNodes(head *ListNode) []int {
	res := make([]int, 0, 10000)
	for head != nil {
		res = append(res, larger(head))
		head = head.Next
	}
	return res
}

func larger(head *ListNode) int {
	val, next := head.Val, head.Next
	for next != nil {
		if val < next.Val {
			return next.Val
		}
		next = next.Next
	}
	return 0
}
