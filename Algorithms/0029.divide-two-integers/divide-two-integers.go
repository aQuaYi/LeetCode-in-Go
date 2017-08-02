package Problem0029

import (
	"math"
)

func divide(dividend int, divisor int) int {
	signD, absD := analysis(dividend)
	signR, absR := analysis(divisor)

	if absD < absR {
		return 0
	}

	res := 1
	temp := absR
	for temp+absR <= absD {
		res++
		temp += absR
	}

	if signD != signR {
		res = res - res - res
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}
func analysis(num int) (sign, abs int) {
	sign = 1
	abs = num
	if num < 0 {
		sign = -1
		abs = abs - abs - abs
	}

	return
}
