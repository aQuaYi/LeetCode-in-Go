package Problem0264

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}

	pos := []int{0, 0, 0}
	candidates := []int{2, 3, 5}
	pNums := []int{2, 3, 5}
	result := make([]int, n)

	result[0] = 1

	for i := 1; i < n; i++ {
		result[i] = minc(candidates)
		for j := 0; j < 3; j++ {
			if result[i] == candidates[j] {
				pos[j]++
				candidates[j] = result[pos[j]] * pNums[j]
			}
		}
	}

	return result[n-1]
}

func minc(candidates []int) int {

	min := candidates[0]
	for i := 1; i < 3; i++ {
		if min > candidates[i] {
			min = candidates[i]
		}
	}
	return min
}
func min3(a, b int) int {
	if a < b {
		return a
	}
	return b
}
