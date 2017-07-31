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

	var cur *ListNode
	head = swap(head)
	cur = head.Next

	for cur != nil && cur.Next != nil {
		cur.Next = swap(cur.Next)
		cur = cur.Next.Next
	}

	return head
}

func swap(l *ListNode) *ListNode {
	if l.Next != nil {
		temp := l.Next
		l.Next = temp.Next
		temp.Next = l
		return temp
	}

	return l
}
