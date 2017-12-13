package Problem0594

func findLHS(nums []int) int {
	r := make(map[int]int, len(nums))
	for _, n := range nums {
		r[n]++
	}

	max := 0
	for n, c := range r {
		t := c + r[n+1]
		if max < t {
			max = t
		}
	}

	return max
}
