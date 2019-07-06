package problem1049

func lastStoneWeightII(stones []int) int {
	hasSum := [3001]bool{}
	hasSum[0] = true
	sum := 0
	for _, s := range stones {
		sum += s
		for i := sum; i >= s; i-- {
			// record all sum of any choose form stones
			hasSum[i] = hasSum[i] || hasSum[i-s]
		}
	}
	// now, sum is sum(stones)
	// take stones to two group，weights are part and sum-part
	// result is sum-part - part
	// best is part = sum/2
	// if not, check part--
	// until part = 0
	part := sum / 2
	for part >= 0 && !hasSum[part] {
		part--
	}
	return sum - part - part
}
