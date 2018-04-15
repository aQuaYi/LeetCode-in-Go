package problem0798

func bestRotation(a []int) int {
	maxScore := -1
	res := -1
	for k := 0; k < len(a); k++ {
		newScore := score(k, a)
		if maxScore < newScore {
			maxScore = newScore
			res = k
		}
	}
	return res
}

func score(k int, a []int) int {
	s := 0
	size := len(a)
	for i, v := range a {
		newIdx := (i - k + size) % size
		if v <= newIdx {
			s++
		}
	}
	return s
}
