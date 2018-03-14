package problem0729

// MyCalendar 是我的日历类
type MyCalendar struct {
	head *event
}
type event struct {
	start, end int
	next       *event
}

// Constructor 构建 MyCanlendar
func Constructor() MyCalendar {
	var head *event
	return MyCalendar{head: head}
}

// Book 可以预约时间
func (m *MyCalendar) Book(start int, end int) bool {
	e := &event{start: start, end: end}

	if m.head == nil || // 还没有事件记录
		e.end <= m.head.start { // 新事件的截止事件，在所有事件开始之前
		e.next, m.head = m.head, e
		return true
	}

	t := &event{next: m.head}
	// 跳过所有在 e 开始之前的事件
	for t.next != nil && t.next.end <= e.start {
		t = t.next
	}

	if t.next == nil || // e 开始事件在所有事件结束之后
		e.end <= t.next.start { // e 结束后，才开始下一个事件
		e.next, t.next = t.next, e
		return true
	}

	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
