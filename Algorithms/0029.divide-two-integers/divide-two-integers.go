package problem0029

import (
	"math"
)

func divide(m, n int) int {
	// 防止有人把0当做除数
	if n == 0 {
		return math.MaxInt32
	}

	signM, absM := analysis(m)
	signN, absN := analysis(n)

	res, _ := d(absM, absN, 1)

	// 修改res的符号
	if signM != signN {
		res = res - res - res
	}

	// 检查溢出
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
		abs = num - num - num
	}

	return
}

// d 计算m/n的值，返回结果和余数
// m >= 0
// n > 0
// count == 1, 代表初始n的个数，在递归过程中，count == 当前的n/初始的n
func d(m, n, count int) (res, remainder int) {
	switch {
	case m < n:
		return 0, m
	case n <= m && m < n+n:
		return count, m - n
	default:
		res, remainder = d(m, n+n, count+count)
		if remainder >= n {
			return res + count, remainder - n
		}

		return
	}
}

// 以下为上述递归方法的普通实现方式
// func d(m, n int) int {
// 	res := 0
// 	rs, ress := []int{n}, []int{1}
// 	temp, i := n+n, 1

// 	for temp <= m {
// 		rs = append(rs, temp)
// 		ress = append(ress, ress[i-1]+ress[i-1])

// 		temp += temp
// 		i++
// 	}

// 	for i := len(rs) - 1; i >= 0; i-- {
// 		if m >= rs[i] {
// 			m -= rs[i]
// 			res += ress[i]
// 		}
// 	}

// 	return res
// }
