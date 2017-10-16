package Problem0234

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type ListNode = kit.ListNode

func isPalindrome(head *ListNode) bool {
	r := reverse(head)

	for head != nil {
		if head.Val != r.Val {
			return false
		}
		head , r = head.Next, r.Next
	}
	return true
}

func reverse(head *ListNode)  *ListNode {
var r *ListNode

for head != nil {
	temp := &ListNode{Val: head.Val}
	if r == nil {
		r = temp
	}else{
		temp.Next = r
		r = temp 
	}
	head = head.Next
}

return r 
}
