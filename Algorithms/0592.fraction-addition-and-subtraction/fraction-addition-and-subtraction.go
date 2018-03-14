package problem0592

import (
	"strconv"
	"strings"
)

func fractionAddition(expression string) string {
	expression = insertPlus(expression)
	es := strings.Split(expression, "+")
	res := getFrac(es[0])
	for i := 1; i < len(es); i++ {
		res = add(res, getFrac(es[i]))
	}
	return res.String()
}

// 在 s[1:] 中的每一个 '-' 前，插入一个 '+'
func insertPlus(s string) string {
	bytes := []byte(s)
	res := make([]byte, 1, len(bytes)*5/4)
	res[0] = bytes[0]
	for i := 1; i < len(bytes); i++ {
		if bytes[i] == '-' {
			res = append(res, '+')
		}
		res = append(res, bytes[i])
	}
	return string(res)
}

type fraction struct {
	num, den int
}

func add(a, b fraction) fraction {
	num := a.num*b.den + b.num*a.den
	den := a.den * b.den
	d := gcd(abs(num), den)
	return fraction{
		num: num / d,
		den: den / d,
	}
}

func getFrac(s string) fraction {
	ss := strings.Split(s, "/")
	n, _ := strconv.Atoi(ss[0])
	d, _ := strconv.Atoi(ss[1])
	return fraction{
		num: n,
		den: d,
	}
}

func (f fraction) String() string {
	return strconv.Itoa(f.num) + "/" + strconv.Itoa(f.den)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
