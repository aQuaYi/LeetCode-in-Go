package Problem0445

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type ListNode = kit.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	s1:=make([]int, 0, 128)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	s2:=make([]int, 0, 128)
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	sum:=0

	list := &ListNode{Val:0}

	for len(s1) > 0 || len(s2)> 0  {

		if len(s1) >0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		if len(s2) >0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		list.Val = sum % 10

		head := & ListNode{Val:sum/ 10}
		head.Next = list
		list = head
		sum/=10
	}
	
	if list.Val == 0 {
		return list.Next
	}
	return list
}

