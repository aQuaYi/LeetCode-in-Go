package Problem0640

import (
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	a, b := analyze(equation)

	if a == 0 {
		if b == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}

	return "x=" + strconv.Itoa(b/a)
}

// 把 equation 整理成 a*x=b
func analyze(equation string) (a, b int) {
	es := strings.Split(equation, "=")
	l, r := es[0], es[1]
	la, lb := sideAnalyze(l)
	ra, rb := sideAnalyze(r)
	return la - ra, rb - lb
}

func sideAnalyze(s string) (a, b int) {
	s = addPlus(s)
	ss := strings.Split(s, "+")
	for _, n := range ss {
		if n[len(n)-1] == 'x' {
			if n == "x" {
				a++
			} else if n == "-x" {
				a--
			} else {
				t, _ := strconv.Atoi(n[:len(n)-1])
				a += t
			}
		} else {
			t, _ := strconv.Atoi(n)
			b += t
		}
	}
	return
}

// 在 s[1:] 中的每个 '-' 前插入 '+'
// 便于分割
func addPlus(s string) string {
	res := make([]byte, 1, len(s)*3/2)
	bs := []byte(s)
	res[0] = bs[0]

	for i := 1; i < len(bs); i++ {
		if bs[i] == '-' {
			res = append(res, '+')
		}
		res = append(res, bs[i])
	}

	return string(res)
}
