package Problem0462

func minMoves2(nums []int) int {
	sum := 0
	size := len(nums)

	for _, n := range nums {
		sum += n
	}

	ave := int(float64(sum)/float64(size) + 0.5)

	res := 0
	for _, n := range nums {
		res += diff(ave, n)
	}

	return res
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
