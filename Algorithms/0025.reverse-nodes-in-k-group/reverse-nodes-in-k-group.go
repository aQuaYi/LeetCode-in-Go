package Problem0025

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {

	return nil
}
func canReverse(head *ListNode, k int) (*ListNode, bool) {
	for head != nil {
		k--
		if k == 0 {
			return head.Next, true
		}
	}
	return nil, false
}

func reverse(head *ListNode) (first, last *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}
	if head.Next.Next == nil {
		last = head
		first = head.Next
		first.Next = last
	}
	var temp *ListNode
	last = head
	first, temp = reverse(head.Next)
	temp.Next = last
	//last.Next = nil
	return first, last
}

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}
