package problem0732

import "sort"

// MyCalendarThree 第二个日历类
type MyCalendarThree struct {
	events [][]int
	k      int
}

// Constructor 返回 MyCalendarTwo
func Constructor() MyCalendarThree {
	return MyCalendarThree{}
}

// Book 可以预约事件，预约成功返回 true
func (m *MyCalendarThree) Book(start, end int) int {
	// e[2] 代表了此时间段的重合程度
	e := []int{start, end, 1}

	n := len(m.events)
	if n == 0 {
		m.events = append(m.events, e)
		m.k = 1
		return m.k
	}

	i := sort.Search(n, func(i int) bool {
		return e[0] < m.events[i][1]
	})
	j := sort.Search(n, func(j int) bool {
		return e[1] <= m.events[j][0]
	})

	// m.events[i:j] 中的事件，都会与 e 有交集
	// 但如果 i==j 的话
	// 说明 e 应该插入到 m.events[i] 的位置
	if i == j {
		m.events = append(m.events, nil)
		copy(m.events[i+1:], m.events[i:])
		m.events[i] = e
		return m.k
	}

	// 更新 m.k
	for k := i; k < j; k++ {
		m.k = max(m.k, m.events[k][2]+1)
	}

	segements := overlap(m.events[i:j], e)
	temp := make([][]int, len(m.events)-j)
	copy(temp, m.events[j:])
	m.events = m.events[:i]
	m.events = append(m.events, segements...)
	m.events = append(m.events, temp...)

	return m.k
}

func overlap(events [][]int, e []int) [][]int {
	res := make([][]int, 0, len(events)*2+1)

	for i := 0; i < len(events); i++ {
		// 切除头部不重叠的部分
		// 运行完成后，
		// e[0]==events[i][0]
		if e[0] < events[i][0] {
			e1 := []int{e[0], events[i][0], 1}
			res = append(res, e1)
			e[0] = events[i][0]
		} else if events[i][0] < e[0] {
			ei := []int{events[i][0], e[0], events[i][2]}
			res = append(res, ei)
			events[i][0] = e[0]
		}
		// 计算其中重叠的部分
		if e[1] < events[i][1] {
			eo := []int{e[0], e[1], events[i][2] + 1}
			ei := []int{e[1], events[i][1], events[i][2]}
			res = append(res, eo, ei)
			e = nil
		} else if e[1] == events[i][1] {
			eo := []int{e[0], e[1], events[i][2] + 1}
			res = append(res, eo)
			e = nil
		} else {
			eo := []int{e[0], events[i][1], events[i][2] + 1}
			res = append(res, eo)
			e[0] = events[i][1]
		}
	}

	if e != nil {
		res = append(res, e)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
