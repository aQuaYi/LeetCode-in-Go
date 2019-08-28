package problem1171

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is ...
type ListNode = kit.ListNode

func removeZeroSumSublists(head *ListNode) *ListNode {
	headPre := &ListNode{
		Val:  2000000, // bigger than 1000*1000
		Next: head,
	}

	node := make(map[int]*ListNode, 1000)
	stack, top := make([]int, 1000), -1
	sum, cur := 0, headPre
	for cur != nil {
		sum += cur.Val
		if node[sum] == nil {
			node[sum] = cur
			top++
			stack[top] = sum
		} else {
			node[sum].Next = cur.Next
			for stack[top] != sum {
				node[stack[top]] = nil
				top--
			}
		}
		cur = cur.Next
	}

	return headPre.Next
}
