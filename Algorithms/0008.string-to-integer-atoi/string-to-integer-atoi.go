package Problem0008

import "strings"
import "math"

func myAtoi(str string) int {
	// 除去 str 首尾的空格
	s := strings.TrimSpace(str)
	if len(s) == 0 {
		return 0
	}

	// 取出 s 的符号和主体 x
	sign, x := getSign(s)

	// 裁剪x丢弃混入的非数字字符
	x = trim(x)

	// 根据sigh和x，转换成int
	return convert(sign, x)
}

func getSign(s string) (int, string) {
	sign := 1
	switch s[0] {
	case '-':
		s = s[1:]
		sign = -1.0
	case '+':
		s = s[1:]
	default:
	}

	return sign, s
}

func trim(s string) string {
	for i := range s {
		if s[i] < '0' || '9' < s[i] {
			return s[:i]
		}
	}

	return s
}

func convert(sign int, s string) int {
	base := 1 * sign
	res := 0
	yes := false

	for i := len(s) - 1; i >= 0; i-- {
		// 为了防止特别长的数字，甚至超过float64的范围，所以，每一步都检查是否溢出
		res, yes = isOverflow(res + (int(s[i])-48)*base)
		if yes {
			return res
		}
		base *= 10
	}

	return res
}

func isOverflow(i int) (int, bool) {
	switch {
	case i > math.MaxInt32:
		return math.MaxInt32, true
	case i < math.MinInt32:
		return math.MinInt32, true
	default:
		return i, false
	}
}
