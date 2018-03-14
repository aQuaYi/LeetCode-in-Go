package problem0086

// ListNode 是链接节点
// 这个不能复制到*_test.go文件中。会导致Travis失败
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 存放 <x 的节点的链
	lessHead := &ListNode{}
	// 存放 >=x 的节点的链
	noLessHead := &ListNode{}
	// Head.Next 才是真正的 head
	// 这样处理是为了 for 循环中的逻辑简单

	lessEnd := lessHead
	noLessEnd := noLessHead

	for head != nil {
		if head.Val < x {
			lessEnd.Next = head
			lessEnd = lessEnd.Next
		} else {
			noLessEnd.Next = head
			noLessEnd = noLessEnd.Next
		}
		head = head.Next
	}

	// 把两部分首尾相连
	lessEnd.Next = noLessHead.Next
	// 注意封闭 noLessEnd
	noLessEnd.Next = nil

	head = lessHead.Next
	return head
}
