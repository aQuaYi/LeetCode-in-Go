package Problem0024

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	t := head.Next
	head.Next = swapPairs(t.Next)
	t.Next = head

	return t
}
