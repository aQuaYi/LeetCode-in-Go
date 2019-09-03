package problem1177

func canMakePaliQueries(s string, queries [][]int) []bool {
	count := [26][]int{}
	for i, letter := range s {
		b := letter - 'a'
		count[b] = append(count[b], i)
	}

	check := func(q []int) bool {
		left, right, k := q[0], q[1], q[2]
		k++
		for i := left; i <= right; i++ {
			// TODO:
			if k < 0 {
				return false
			}
		}
		return true
	}

	return nil
}
