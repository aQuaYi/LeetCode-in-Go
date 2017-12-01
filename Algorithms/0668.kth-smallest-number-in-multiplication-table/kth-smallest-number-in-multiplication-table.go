package Problem0668

func findKthNumber(m int, n int, k int) int {
	mn := m * n
	for i := 1; i <= mn; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 {
				k--
				if k == 0 {
					return i
				}
			}
		}
	}

	return mn
}
