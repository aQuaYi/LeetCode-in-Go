package problem0640

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
	// la*x+lb=ra*x+rb
	la, lb := sideAnalyze(l)
	ra, rb := sideAnalyze(r)
	return la - ra, rb - lb
}

// 	把 a1*x+b1+a2*x+b2 整理成 a*x+b
func sideAnalyze(s string) (a, b int) {
	s = s[:1] + strings.Replace(s[1:], "-", "+-", -1)
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
