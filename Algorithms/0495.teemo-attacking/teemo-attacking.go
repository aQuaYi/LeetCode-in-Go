package problem0495

func findPoisonedDuration(time []int, duration int) int {
	res := 0
	n := len(time)
	if n == 0 {
		return 0
	}

	for i := 0; i+1 < n; i++ {
		res += min(duration, time[i+1]-time[i])
	}

	return res + duration
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
