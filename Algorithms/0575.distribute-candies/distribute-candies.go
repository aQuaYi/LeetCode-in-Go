package problem0575

func distributeCandies(candies []int) int {
	n := len(candies)

	r := make(map[int]bool, n)
	for _, c := range candies {
		r[c] = true
	}

	return min(len(r), n/2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
