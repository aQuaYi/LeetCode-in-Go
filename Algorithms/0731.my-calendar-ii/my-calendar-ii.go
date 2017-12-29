package Problem0731

// MyCalendarTwo 第二个日历类
type MyCalendarTwo struct {
	head, tail *event
}

type event struct {
	start, end int
	count      int
	pre, next  *event
}

// Constructor 返回 MyCalendarTwo
func Constructor() MyCalendarTwo {
	head := &event{start: -1, end: -1}
	tail := &event{start: 1e9 + 1, end: 1e9 + 1}
	head.next = tail
	tail.pre = head
	return MyCalendarTwo{head: head, tail: tail}
}

// Book 可以预约事件，预约成功返回 true
func (m *MyCalendarTwo) Book(start, end int) bool {
	e := &event{start: start, end: end, count: 1}

	cur := m.head
	for cur.end <= e.start {
		cur = cur.next
	}

	for e != nil {
		if e.end <= cur.start {
			e.pre, e.next = cur.pre, cur
			e.pre.next, e.next.pre = e, e
			e = nil
			continue
		}

	}

	return false
}

// 在 a 的 pre 方向插入 aPre
func addPre(a, aPre *event) {
	aPre.next, aPre.pre = a, a.pre
	aPre.pre.next, aPre.next.pre = aPre, aPre
}

func split(e, c *event) (ecp, ec, ecn *event) {

}

func isAddable(headPre, tailNext, e *event) bool {
	cur := headPre.next
	for cur != tailNext {
		if cur.count >= 2 {
			return false
		}
	}

	return true
}

// if old.end<new.end 返回  &event{start:old.end,end:new.end}
// 否则，返回 nil
func overlap(old, new *event) *event {
	if new.start < old.start {
		t := &event{start: new.start, end: old.start, count: new.count}
		old.pre.next = t
		t.next = old
		t.pre = old.pre
		old.pre = t
	} else {

	}

	return nil
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
