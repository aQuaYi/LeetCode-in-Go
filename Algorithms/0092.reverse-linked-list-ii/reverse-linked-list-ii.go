package problem0092

// ListNode is Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	// 添加 headPre 是为了使得 m >= 2
	// 避免了 reverse 的部分包含 head
	// 让后面 split 和组合的逻辑简单清晰
	headPre := &ListNode{}
	headPre.Next = head
	m++
	n++

	// 按照 m，n 的值，把链切断
	mPre, mNode, nNext := split(headPre, m, n)
	// reverse 中间段
	h, e := reverse(mNode)
	// 把链条接好
	mPre.Next = h
	e.Next = nNext
	// headPre.Next 才是真正的 head
	return headPre.Next
}

func split(head *ListNode, m, n int) (mPre, mNode, nNext *ListNode) {
	i := 1
	for head != nil {
		if i == m-1 {
			mPre = head
			mNode = head.Next
		}
		if i == n {
			nNext = head.Next
			// head.Next = nil，很重要
			// 不做的话，reverse会出错
			head.Next = nil
			break
		}
		head = head.Next
		i++
	}

	return
}

// 返回新链条的 head 和 end
func reverse(head *ListNode) (h, e *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	var end *ListNode

	h, end = reverse(head.Next)
	end.Next = head
	e = head

	return
}
