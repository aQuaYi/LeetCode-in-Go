package problem1217

func minCostToMoveChips(chips []int) int {
	odd, even := 0, 0
	for _, p := range chips {
		if p%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return min(odd, even)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
