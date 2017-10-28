package Problem0373

func kSmallestPairs(a, b []int, k int) [][]int {
	res := make([][]int, 0, k)
	m, n := len(a), len(b)
	if m == 0 || n == 0 || k == 0 {
		return res
	}

	res = append(res, []int{a[0], b[0]})
	k--
	var i, j int

	for i < m && j < n {
		if k == 0 {
			return res
		}

		if i == m-1 {
			if j+1 < n {
				res = append(res, []int{a[i], b[j+1]})
			}
			j++

			continue
		}
		if i+1 < m && j+1 < n && a[i+1]+b[j] < a[i]+b[j+1] {

		}
	}

	return res
}
