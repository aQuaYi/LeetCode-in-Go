package problem0817

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * 	return 0
}
*/

// ListNode is from kit
type ListNode = kit.ListNode

func numComponents(head *ListNode, G []int) int {
	isInG := [10001]bool{}
	for i := range G {
		isInG[G[i]] = true
	}

	res := 0
	for head != nil {
		if isInG[head.Val] &&
			(head.Next == nil || !isInG[head.Next.Val]) {
			res++
		}
		head = head.Next
	}

	return res
}
