package problem0902

import (
	"strconv"
)

func atMostNGivenDigitSet(D []string, N int) int {
	ns := strconv.Itoa(N)
	size := len(D)
	res := zeroLead(size, len(ns))
	return res + equal(D, ns)
}

// 从 N 的高位往低位防止 D 中的数字，分为 3 种情况
// 1. 可以选择不放置数字，但是一旦放置了，后续低位就不能不放置了，这一类数的特点是以 0 开头，其个数使用 zeroLead 函数计算。
// 2. 比 N 对应位上的数字小，其后续低位可以任意放置 D 中的数字，其个数使用 less 函数计算。
// 3. 与 N 对应位上的数字一样大，其后续低位上的数字需要分情况讨论，这一类数的个数，使用 equal 函数，递归计算。

func equal(D []string, ns string) int {
	if len(ns) == 0 {
		return 1
	}
	size := len(D)
	res := 0
	head := ns[0:1]
	for _, d := range D {
		switch {
		case d < head:
			res += less(size, len(ns))
		case d == head:
			res += equal(D, ns[1:])
		}
	}
	return res
}

// size 个数，组合成低于 length 位的数字的组合数
func zeroLead(size, length int) int {
	p := size
	res := 0
	for length > 1 {
		res += p
		p *= size
		length--
	}
	return res
}

// size 个数，组成 length-1 位数字的组合数
func less(size, length int) int {
	res := 1
	for length > 1 {
		res *= size
		length--
	}
	return res
}
