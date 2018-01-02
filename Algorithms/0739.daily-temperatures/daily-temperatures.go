package Problem0739

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		j := i + 1
		for j < n && temperatures[i] >= temperatures[j] {
			j++
		}

		if j < n {
			res[i] = j - i
		}
	}

	return res
}
