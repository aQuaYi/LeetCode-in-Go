package Problem0365

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}

	if x > y {
		return canMeasureWater(y, x, z)
	}

	if x == 0 {
		return y == z
	}

	if z > x+y {
		return false
	}

	if z == x+y || z%x == 0 || z == y {
		return true
	}

	for i := 2; i < y; i++ {
		if x*i%y == z {
			return true
		}
	}

	return false
}
