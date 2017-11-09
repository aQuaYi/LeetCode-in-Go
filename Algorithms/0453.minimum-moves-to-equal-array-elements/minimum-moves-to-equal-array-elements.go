package Problem0453

func minMoves(nums []int) int {
	sum, min := 0, 1<<31-1
	for _, n := range nums {
		sum += n
		if min > n {
			min = n
		}
	}

	return sum - min*len(nums)
}
