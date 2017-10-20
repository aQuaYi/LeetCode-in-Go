package Problem0365

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}

	if x > y {
		// 保证 x <= y
		return canMeasureWater(y, x, z)
	}

	if x == 0 {
		return y == z
	}

	a, b := y, x
	for a%b != 0 {
		b, a = a%b, b
	}

	return z%b == 0 && x+y >= z
}
