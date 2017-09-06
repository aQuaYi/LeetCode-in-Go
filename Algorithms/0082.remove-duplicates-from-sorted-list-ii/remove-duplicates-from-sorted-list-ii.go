package Problem0082

// ListNode is singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 长度 <=1 的 list ，可以直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 处理 head 重复情况
	// 上面刚刚检查过了，head != nil && head.Next != nil
	if head.Val == head.Next.Val {
		val := head.Val
		head = head.Next.Next

		for head != nil && head.Val == val {
			head = head.Next
		}

		// 值为 val 的 node，已经全部删除了
		return deleteDuplicates(head)
	}

	// 处理 head 后面元素出现重复的情况
	head.Next = deleteDuplicates(head.Next)

	return head
}
