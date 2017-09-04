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

	end := head
	// size 是 List 的长度
	size := 1
	for end.Next != nil {
		end = end.Next
		size++
	}

	end.Next = head

	// steps 是 head 需要移动的步数
	// k%size 因为 k 很可能大于 size
	steps := size - k%size

	for steps > 0 {
		head = head.Next
		end = end.Next
		steps--
	}

	end.Next = nil

	return head
}
