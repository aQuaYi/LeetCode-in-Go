package Problem0659

func isPossible(nums []int) bool {
	count := make(map[int]int, 10000)
	next := make(map[int]int, 10000)
	for i := range nums {
		count[nums[i]]++
	}

	for _, n := range nums {
		if count[n] == 0 {
			continue
		}

		if next[n] > 0 {
			next[n]--
			next[n+1]++
		} else if count[n+1] > 0 && count[n+2] > 0 {
			count[n+1]--
			count[n+2]--
			next[n+3]++
		} else {
			return false
		}
		count[n]--
	}

	return true
}
