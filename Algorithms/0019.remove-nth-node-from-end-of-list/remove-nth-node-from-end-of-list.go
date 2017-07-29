package Problem0019

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 是节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nd, n := getNDaddy(head, n)
	if nd == head && n == 0 {
		return head.Next
	}
	nd.Next = nd.Next.Next
	return head
}

func getNDaddy(head *ListNode, n int) (*ListNode, int) {
	NDaddy := head
	for head != nil {
		if n < 0 {
			NDaddy = NDaddy.Next
		}
		n--
		head = head.Next
	}
	return NDaddy, n
}
