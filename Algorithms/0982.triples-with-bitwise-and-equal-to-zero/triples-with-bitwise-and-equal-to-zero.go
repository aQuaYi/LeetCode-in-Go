package problem0982

func countTriplets(A []int) int {
	count := [1 << 16]int{}
	for _, Ai := range A {
		for _, Aj := range A {
			count[Ai&Aj]++
		}
	}

	res := 0
	for AiAndAj, c := range count {
		if c == 0 {
			continue
		}
		for _, Ak := range A {
			if AiAndAj&Ak == 0 {
				res += c
			}
		}
	}

	return res
}
