package problem0365

func canMeasureWater(x int, y int, z int) bool {
	// 不需要水壶就能实现
	if z == 0 {
		return true
	}

	// 两个水壶都不够装
	if x+y < z {
		return false
	}

	// 此段以后，需要 x <= y
	if x > y {
		x, y = y, x
	}

	// 只有一个水壶的话，
	// 只能提供另一个水壶的容量
	if x == 0 {
		return y == z
	}

	// y%x == 0 时，
	// x 是初始 x 和 y 的最大公约数
	// z%x == 0 的容量都可以提供
	for y%x != 0 {
		x, y = y%x, x
	}

	return z%x == 0
}
