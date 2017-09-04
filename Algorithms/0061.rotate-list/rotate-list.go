package Problem0061

// ListNode 是题目预定义的节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	fast := head
	for i := 0; i < k; i++ {
		if fast.Next == nil {
			// 处理 k 大于 list 的长度的情况
			// i+1 就是 list 的长度
			// 这一步很巧妙
			return rotateRight(head, k%(i+1))
		}
		fast = fast.Next
	}

	slow := head
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	newHead := slow.Next
	slow.Next, fast.Next = nil, head

	return newHead
}
