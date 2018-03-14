package problem0025

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k < 2 {
		return head
	}

	next, ok := needReverse(head, k)
	if ok {
		head, tail := reverse(head)
		// 递归
		// 把整理好了的前k个节点的尾部，指向整理好了的后面节点的head
		tail.Next = reverseKGroup(next, k)
		return head
	}

	return head
}

// 判断是否有前k个节点需要逆转。
// 需要的话
// 会把KthNode.Next = nil，把k和k+1节点斩断，便于前k个节点的逆转。
func needReverse(head *ListNode, k int) (begin *ListNode, ok bool) {
	for head != nil {
		if k == 1 {
			begin = head.Next
			// 把前k与后面的节点斩断, 便于reverse
			head.Next = nil
			return begin, true
		}

		head = head.Next
		k--
	}

	return nil, false
}

// 返回逆转后的首尾节点
func reverse(head *ListNode) (first, last *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	gotLast := false

	for head != nil {
		temp := head.Next
		head.Next = first
		first = head
		head = temp

		if !gotLast {
			last = first
			gotLast = true
		}
	}

	return first, last
}

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}
