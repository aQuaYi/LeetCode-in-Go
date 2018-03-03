package Problem0765

func minSwapsCouples(row []int) int {
	N := len(row) / 2

	uncoupleCount := 0
	for i := 0; i < 2*N; i += 2 {
		if !isCouple(row[i], row[i+1]) {
			uncoupleCount++
		}
	}

	return max(0, uncoupleCount-1)
}

func isCouple(a, b int) bool {
	if a > b {
		a, b = b, a
	}

	return a+1 == b && a%2 == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
