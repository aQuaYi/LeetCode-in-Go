package problem0822

func flipgame(fronts []int, backs []int) int {
	count := make(map[int]int, len(fronts)+len(backs))

	for i := range fronts {
		count[fronts[i]]++
	}

	for i := range backs {
		count[backs[i]]++
	}

	upLimit := 2001
	min := upLimit
	for n, c := range count {
		if c == 1 && min > n {
			min = n
		}
	}

	if min == upLimit {
		return 0
	}
	return min

}
