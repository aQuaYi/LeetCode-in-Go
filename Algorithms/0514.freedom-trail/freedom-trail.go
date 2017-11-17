package Problem0514

func findRotateSteps(ring string, key string) int {
	m := make(map[rune][]int, 26)
	for i, r := range ring {
		m[r] = append(m[r], i)
	}

	ss := make([]status, 1, len(ring))
	ss[0] = status{
		index: 0,
		steps: 0,
	}

	size := len(ring)

	for _, r := range key {
		temp := make([]status, len(m[r]))
		for i, d := range m[r] {
			s := fromTo(ss[0], d, size)
			for i := 1; i < len(ss); i++ {
				s.steps = min(s.steps, fromTo(ss[i], d, size).steps)
			}
			temp[i] = s
		}
		ss = temp
	}

	res := ss[0].steps
	for i := 1; i < len(ss); i++ {
		res = min(res, ss[i].steps)
	}

	return res
}

// status 记录了
type status struct {
	index, steps int
}

func fromTo(s status, d, size int) status {
	a := min(s.index, d)
	b := max(s.index, d)

	return status{
		index: d,
		steps: s.steps + min(b-a, size-(b-a)) + 1,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
