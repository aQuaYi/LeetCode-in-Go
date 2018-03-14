package problem0319

func bulbSwitch(n int) int {
	return intSqrt(n)
}

// intSqrt 返回 x 的平方根的整数部分
func intSqrt(x int) int {
	root := x

	for root*root > x {
		root = (root + x/root) / 2
	}

	return root
}
