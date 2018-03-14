package problem0021

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 有一条链为nil，直接返回另一条链
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 此时，两条链都不为nil，可以直接使用l.Val，而不用担心panic
	// 在l1和l2之间，选择较小的节点作为head，并设置好node
	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	// 循环比较l1和l2的值，始终选择较小的值连上node
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}

		// 有了这一步，head才是一个完整的链
		node = node.Next
	}

	if l1 != nil {
		// 连上l1剩余的链
		node.Next = l1
	}

	if l2 != nil {
		// 连上l2剩余的链
		node.Next = l2
	}

	return head
}
