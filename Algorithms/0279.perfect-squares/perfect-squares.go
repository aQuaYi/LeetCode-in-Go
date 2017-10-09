package Problem0279

func numSquares(n int) int {
	if n < 4 {
		return n
	}

	max := intSqrt(n)
	min := intSqrt(max) + 1
	res := n
	var i2, temp int

	for i := min; i <= max; i++ {
		i2 = i * i
		temp = n/i2 + numSquares(n%i2)
		if res > temp {
			res = temp
		}
	}

	return res
}

// 返回 x 的平方根的整数部分
// 这个函数比 int(math.Sqrt(float64(x))) 快的多
// 详见 benchmark 的结果
func intSqrt(x int) int {
	res := x

	// 牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}
