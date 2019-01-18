package problem0008

import "strings"
import "math"

func myAtoi(s string) int {
	return convert(clean(s))
}

func clean(s string) (sign int, abs string) {
	// 除去首尾的空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}

	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
	default:
		abs = ""
		return
	}

	for i, b := range abs {
		if b < '0' || '9' < b {
			abs = abs[:i]
			break
		}
	}

	return
}

func convert(sign int, absStr string) int {
	abs := 0

	for _, b := range absStr {
		abs = abs*10 + int(b-'0')
		// 检查溢出
		switch {
		case sign == 1 && abs > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && sign*abs < math.MinInt32:
			return math.MinInt32
		}
	}

	return sign * abs
}
