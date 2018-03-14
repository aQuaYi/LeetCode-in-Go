package problem0206

func reverseList(head *ListNode) *ListNode {
	// prev 是所有已经逆转的节点的head
	var prev *ListNode

	// head 是下一个被逆转的节点
	for head != nil {
		// 让temp指向head.Next, 免得head.Next不见了.
		temp := head.Next
		// head称为已经逆转的节点的新head
		head.Next = prev
		// 让prev重新称为所有已被逆转节点的head
		prev = head
		// 让head指向下一个被逆转的节点
		head = temp
	}

	return prev
}

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}
