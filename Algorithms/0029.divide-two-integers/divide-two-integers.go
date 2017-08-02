package Problem0029

import (
	"math"
)

func divide(m, n int) int {
	signM, absM := analysis(m)
	signN, absN := analysis(n)

	if absM < absN {
		return 0
	}

	res := d(absM, absN)

	if signM != signN {
		res = res - res - res
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

func d(m, n int) int {
	res := 0
	rs, ress := []int{n}, []int{1}
	temp, i := n+n, 1

	for temp <= m {
		rs = append(rs, temp)
		ress = append(ress, ress[i-1]+ress[i-1])

		temp += temp
		i++
	}

	for i := len(rs) - 1; i >= 0; i-- {
		if m >= rs[i] {
			m -= rs[i]
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
		abs = num - num - num
	}

	return
}
