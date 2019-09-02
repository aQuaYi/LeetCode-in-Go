package problem1176

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	point := func(t int) int {
		if t < lower {
			return -1
		}
		if t > upper {
			return 1
		}
		return 0
	}

	T := 0
	res := 0
	for i, c := range calories {
		T += c
		if i < k-1 {
			continue
		}
		if i-k >= 0 {
			T -= calories[i-k]
		}
		res += point(T)
	}

	return res
}
