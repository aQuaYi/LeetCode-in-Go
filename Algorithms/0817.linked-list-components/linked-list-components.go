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
	isInG := make(map[int]bool, len(G))
	for i := range G {
		isInG[G[i]] = true
	}

	res := 0

	for head != nil {
		/**
		 * head: 0->1->2->3->4
		 * G = [0, 3, 1, 4]
		 * 结果为 2
		 * 可以按照 connected components 划分 head
		 * head: (0->1)->2->(3->4)
		 * 每个单独的 connected components 的特点是
		 * 最后一个 node 的 Next.Val 不在 G 中，或者是 list 的结尾
		 * 那么，统计下面 if 成立的次数，就是所需的结果了
		 */
		if isInG[head.Val] &&
			(head.Next == nil || !isInG[head.Next.Val]) {
			res++
		}
		head = head.Next
	}

	return res
}
