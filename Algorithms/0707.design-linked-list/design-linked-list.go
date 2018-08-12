package problem0707

// MyLinkedList object will be instantiated and called as such:
// obj := Constructor();
// param_1 := obj.Get(index);
// obj.AddAtHead(val);
// obj.AddAtTail(val);
// obj.AddAtIndex(index,val);
// obj.DeleteAtIndex(index);
type MyLinkedList struct {
	size       int
	head, tail *node
}

type node struct {
	val  int
	next *node
}

// Constructor initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *MyLinkedList) Get(index int) int {
	switch {
	case index < 0 || l.size <= index:
		return -1
	case index == 0:
		return l.head.val
	case index == l.size-1:
		return l.tail.val
	}

	i, cur := 0, l.head
	for i < index {
		cur = cur.next
		i++
	}
	return cur.val
}

// AddAtHead a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (l *MyLinkedList) AddAtHead(val int) {
	nd := &node{val: val}
	if l.size == 0 {
		l.head = nd
		l.tail = nd
	} else {
		nd.next = l.head
		l.head = nd
	}
	l.size++
}

// AddAtTail append a node of value val to the last element of the linked list.
func (l *MyLinkedList) AddAtTail(val int) {
	nd := &node{val: val}
	if l.size == 0 {
		l.head = nd
		l.tail = nd
	} else {
		l.tail.next = nd
		l.tail = nd
	}
	l.size++
}

// AddAtIndex add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	switch {
	case index == 0:
		l.AddAtHead(val)
		return
	case index == l.size:
		l.AddAtTail(val)
		return
	case index < 0 || l.size < index:
		return
	}

	nd := &node{val: val}

	i, cur := 0, l.head
	for i+1 < index {
		cur = cur.next
		i++
	}

	nd.next = cur.next
	cur.next = nd

	l.size++
}

// DeleteAtIndex delete the index-th node in the linked list, if the index is valid.
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if l.size == 1 {
			l.head, l.tail = nil, nil
		} else {
			l.head = l.head.next
		}
		l.size--
		return
	}

	if index < 0 || l.size <= index {
		return
	}

	i, cur := 0, l.head
	for i+1 < index {
		cur = cur.next
		i++
	}

	cur.next = cur.next.next

	l.size--
}
