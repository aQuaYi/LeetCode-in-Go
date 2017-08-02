package Problem0029

import (
	"math"
)

func divide(dividendD int, divisorR int) int {
	signD, absD := analysis(dividendD)
	signR, absR := analysis(divisorR)

	if absD < absR {
		return 0
	}

	res := 1
	if absR == 1 {
		res = absD
	} else {
		res = d(absD, absR)
	}

	if signD != signR {
		res = res - res - res
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

func d(D, R int) int {
	res := 0
	rs, ress := []int{R}, []int{1}
	temp, i := R+R, 1

	for temp <= D {
		rs = append(rs, temp)
		ress = append(ress, ress[i-1]+ress[i-1])

		temp += temp
		i++
	}

	for i := len(rs) - 1; i >= 0; i-- {
		if D >= rs[i] {
			D -= rs[i]
			res += ress[i]
		}
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
