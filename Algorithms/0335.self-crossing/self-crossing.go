package Problem0335

func isSelfCrossing(x []int) bool {
	size := len(x)

	cross4Line := func(i int) bool {
		return i >= 3 && x[i] >= x[i-2] && x[i-1] <= x[i-3]
	}

	cross5Line := func(i int) bool {
		return i >= 4 && x[i-1] == x[i-3] && x[i]+x[i-4] >= x[i-2]
	}

	cross6Line := func(i int) bool {
		return i >= 5 && x[i-2] >= x[i-4] && x[i]+x[i-4] >= x[i-2] && x[i-1] <= x[i-3] && x[i-1]+x[i-5] >= x[i-3]
	}

	var i int
	for i = 3; i < size; i++ {
		if cross4Line(i) || cross5Line(i) || cross6Line(i) {
			return true
		}
	}

	return false
}
