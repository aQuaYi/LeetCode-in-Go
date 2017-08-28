package Problem0229

func majorityElement(a []int) []int {
	l := len(a)
	if l < 2 {
		return a
	}

	res := []int{}
	rec := make(map[int]int, l)

	for i := 0; i < l; i++ {
		rec[a[i]]++
	}

	for k, n := range rec {
		if 3*n > l {
			res = append(res, k)
		}
	}

	return res
}
