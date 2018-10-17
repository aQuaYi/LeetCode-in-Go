package problem0061

// ListNode 是题目预定义的节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	tail := head
	for i := 0; i < k; i++ {
		if tail.Next == nil {
			// 处理 k 大于 list 的长度的情况
			// i+1 就是 list 的长度
			return rotateRight(head, k%(i+1))
		}
		tail = tail.Next
	}

	newTail := head
	for tail.Next != nil {
		newTail, tail = newTail.Next, tail.Next
	}

	newHead := newTail.Next
	newTail.Next, tail.Next = nil, head

	return newHead
}
