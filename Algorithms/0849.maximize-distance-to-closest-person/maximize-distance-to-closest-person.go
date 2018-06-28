package problem0849

func maxDistToClosest(seats []int) int {
	size := len(seats)
	maxDis := 0
	// e 代表了连续空位的个数
	// 当连续空位两边都有人的时候，maxDis = (e+e%2)/2
	// 如果有一边没人的话，      maxDis = e
	e := 0
	for i := 0; i < size; i++ {
		if e == i {
			// 说明 seats[0:i] 全是 0
			maxDis = e
		} else {
			maxDis = max(maxDis, (e+e%2)/2)
		}
		if seats[i] == 1 {
			e = 0
		} else {
			e++
		}
	}

	// 当 seats[size-1]==0 的时候
	// e 最后的值，有可能 > maxDis
	return max(maxDis, e)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
