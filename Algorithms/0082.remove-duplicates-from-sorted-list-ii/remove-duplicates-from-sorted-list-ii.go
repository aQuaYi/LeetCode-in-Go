package Problem0082

// ListNode is singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next != nil && head.Val == head.Next.Val {
		val := head.Val
		head = head.Next.Next

		for head != nil && head.Val == val {
			head = head.Next
		}

		return deleteDuplicates(head)
	}

	head.Next = deleteDuplicates(head.Next)

	return head
}
