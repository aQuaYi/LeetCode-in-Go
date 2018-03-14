package problem0335

func isSelfCrossing(x []int) bool {
	size := len(x)

	is4th := func(i int) bool {
		return i+3 < size && x[i+2] <= x[i] && x[i+3] >= x[i+1]
	}

	is5th := func(i int) bool {
		return i+4 < size && x[i+3] == x[i+1] && x[i+4]+x[i] >= x[i+2]
	}

	is6th := func(i int) bool {
		return i+5 < size && x[i+2] > x[i] && x[i+3] > x[i+1] && x[i+4] <= x[i+2] && x[i+4]+x[i] >= x[i+2] && x[i+5]+x[i+1] >= x[i+3]
	}

	for i := 0; i < size; i++ {
		if is4th(i) || is5th(i) || is6th(i) {
			return true
		}
	}

	return false
}
