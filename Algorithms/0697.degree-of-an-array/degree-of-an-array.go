package problem0697

func findShortestSubArray(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}

	first := make(map[int]int, size)
	count := make(map[int]int, size)
	maxCount := 1
	minLen := size
	for i, n := range nums {
		count[n]++
		if count[n] == 1 {
			first[n] = i
		} else {
			l := i - first[n] + 1
			if maxCount < count[n] ||
				(maxCount == count[n] && minLen > l) {
				maxCount = count[n]
				minLen = l
			}
		}
	}

	if len(count) == size {
		return 1
	}
	return minLen
}
