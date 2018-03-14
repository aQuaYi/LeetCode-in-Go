package problem0024

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 让temp指向head.Next节点
	temp := head.Next
	// 让head.Next指向转换好了temp.Next节点
	head.Next = swapPairs(temp.Next)
	// 让temp.Next指向head节点
	temp.Next = head
	// temp成为新的head节点

	return temp
}
