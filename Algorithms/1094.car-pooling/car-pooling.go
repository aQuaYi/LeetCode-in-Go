package problem1094

func carPooling(trips [][]int, capacity int) bool {
	location := [1001]int{}
	for _, t := range trips {
		num, start, end := t[0], t[1], t[2]
		location[start] += num
		location[end] -= num
	}

	p := 0
	for _, c := range location {
		p += c
		if p > capacity {
			return false
		}
	}
	return true
}
