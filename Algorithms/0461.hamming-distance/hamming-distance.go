package Problem0461

func hammingDistance(x int, y int) int {
	x ^= y

	res := 0
	for x > 0 {
		if x&1 == 1 {
			res++
		}
		x >>= 1
	}

	return res
}
