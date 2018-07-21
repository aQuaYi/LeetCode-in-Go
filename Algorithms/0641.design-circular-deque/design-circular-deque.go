package problem0641

// MyCircularDeque 结构体
type MyCircularDeque struct {
	f, r     *node
	len, cap int
}
type node struct {
	value     int
	pre, next *node
}

// Constructor initialize your data structure here. Set the size of the deque to be k.
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cap: k,
	}
}

// InsertFront adds an item at the front of Deque. Return true if the operation is successful.
func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.len == d.cap {
		return false
	}
	n := &node{
		value: value,
	}
	if d.len == 0 {
		d.f = n
		d.r = n
	} else {
		n.next = d.f
		d.f.pre = n
		d.f = n
	}
	d.len++
	return true
}

// InsertLast adds an item at the rear of Deque. Return true if the operation is successful.
func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.len == d.cap {
		return false
	}
	n := &node{
		value: value,
	}
	if d.len == 0 {
		d.f = n
		d.r = n
	} else {
		n.pre = d.r
		d.r.next = n
		d.r = n
	}
	d.len++
	return true
}

// DeleteFront deletes an item from the front of Deque. Return true if the operation is successful.
func (d *MyCircularDeque) DeleteFront() bool {
	if d.len == 0 {
		return false
	}
	if d.len == 1 {
		d.f, d.r = nil, nil
	} else {
		d.f = d.f.next
		d.f.pre = nil
	}
	d.len--
	return true
}

// DeleteLast deletes an item from the rear of Deque. Return true if the operation is successful.
func (d *MyCircularDeque) DeleteLast() bool {
	if d.len == 0 {
		return false
	}
	if d.len == 1 {
		d.f, d.r = nil, nil
	} else {
		d.r = d.r.pre
		d.r.next = nil
	}
	d.len--
	return true
}

// GetFront get the front item from the deque.
func (d *MyCircularDeque) GetFront() int {
	if d.len == 0 {
		return -1
	}
	return d.f.value
}

// GetRear get the last item from the deque.
func (d *MyCircularDeque) GetRear() int {
	if d.len == 0 {
		return -1
	}
	return d.r.value
}

// IsEmpty checks whether the circular deque is empty or not.
func (d *MyCircularDeque) IsEmpty() bool {
	return d.len == 0
}

// IsFull checks whether the circular deque is full or not.
func (d *MyCircularDeque) IsFull() bool {
	return d.len == d.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
