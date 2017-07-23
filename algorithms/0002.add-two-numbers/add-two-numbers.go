package Problem0002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	r := result
	v, n := 0, 0

	for {
		v, n = add(l1, l2, n)
		r.Val = v

		l1 = next(l1)
		l2 = next(l2)

		if !(l1 != nil || l2 != nil) {
			break
		}

		r.Next = &ListNode{}
		r = r.Next
	}

	if n == 1 {
		r.Next = &ListNode{Val: n}
	}

	return result
}

func next(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return nil
}

func add(n1, n2 *ListNode, i int) (v, n int) {
	if n1 != nil {
		v += n1.Val
	}

	if n2 != nil {
		v += n2.Val
	}

	v += i
	if v > 9 {
		v -= 10
		n = 1
	}
	return
}
