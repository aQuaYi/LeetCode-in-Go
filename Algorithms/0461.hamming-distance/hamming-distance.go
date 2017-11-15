package Problem0461

func hammingDistance(x int, y int) int {
	if x > y {
		x, y = y, x
	}

	res := 0
	for x < y {
		x <<= 1
		res++
	}

	return res
}
