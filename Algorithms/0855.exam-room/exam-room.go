package problem0855

// ExamRoom 是题目需要的结构体
type ExamRoom struct {
	seats *[]int
	N     int
}

// Constructor 创建 ExamRoom
func Constructor(N int) ExamRoom { // 如果返回 *ExamRoom 后面可以写的好看一些
	t := make([]int, 0, 10000)
	return ExamRoom{
		seats: &t,
		N:     N,
	}
}

// Seat 表示有人入座
func (r *ExamRoom) Seat() int {
	if len(*r.seats) == 0 {
		*r.seats = append(*r.seats, 0)
		return 0
	}

	f := -(*r.seats)[0]
	*r.seats = append(*r.seats, 2*r.N-2-(*r.seats)[len(*r.seats)-1])
	idx := -1
	dis := -1
	seat := -1

	for i := 0; i < len(*r.seats); i++ {
		newDis := ((*r.seats)[i] - f) / 2
		if dis < newDis {
			idx = i
			dis = newDis
			seat = f + dis
		}
		f = (*r.seats)[i]
	}

	copy((*r.seats)[idx+1:], (*r.seats)[idx:])

	(*r.seats)[idx] = seat

	return seat
}

// Leave 表示有人离座
func (r *ExamRoom) Leave(p int) {
	size := len(*r.seats)
	left, right := 0, size-1

	for { // 题目保证了 p 一定在 r.seats 中
		m := (left + right) / 2
		mp := (*r.seats)[m]
		switch {
		case mp < p:
			left = m + 1
		case p < mp:
			right = m - 1
		default:
			copy((*r.seats)[m:], (*r.seats)[m+1:])
			*r.seats = (*r.seats)[:size-1]
			return
		}
	}
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(N);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
