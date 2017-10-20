package Problem0365

func canMeasureWater(x int, y int, z int) bool {
	if x > y {
		return canMeasureWater(y, x, z)
	}

	if z > x+y {
		return false
	}

	if z%x == 0 || z == y || z == x+y {
		return true
	}

	for i := 2; i < y; i++ {
		if x*i%y == z {
			return true
		}
	}

	return false
}
