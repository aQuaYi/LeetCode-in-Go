package problem0726

import (
	"sort"
	"strconv"
)

func countOfAtoms(formula string) string {
	return parse(count(formula))
}

func count(formula string) map[string]int {
	rec := make(map[string]int, len(formula)/2)
	var update = func(newRec map[string]int, times int) {
		for atom, c := range newRec {
			rec[atom] += c * times
		}
	}

	atoms := ""
	for len(formula) > 0 {
		atoms, formula = cut(formula)
		if atoms[0] == '(' {
			newFormula, times := dealParenthese(atoms)
			newRec := count(newFormula)
			update(newRec, times)
		} else {
			atom, num := getAtomAndNum(atoms)
			rec[atom] += num
		}
	}

	return rec
}

func cut(formula string) (string, string) {
	i := 1
	if formula[0] == '(' {
		i = jump(formula)
	}
	for i < len(formula) &&
		!isUpper(formula[i]) &&
		formula[i] != '(' {
		i++
	}
	return formula[:i], formula[i:]
}

func dealParenthese(s string) (string, int) {
	num, i := getNum(s)
	return s[1 : i-1], num
}

func getAtomAndNum(s string) (string, int) {
	num, i := getNum(s)
	return s[:i], num
}

// 对于 "Ab321" 返回， 321 和 '3' 的索引号 2
func getNum(s string) (num, i int) {
	i = len(s)
	for 0 <= i-1 && isNum(s[i-1]) {
		i--
	}
	num = 1
	if i == len(s) {
		return
	}
	num, _ = strconv.Atoi(s[i:])
	return
}

func isNum(b byte) bool {
	return '0' <= b && b <= '9'
}

func isUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

// jump 跳过了圆括号部分
func jump(s string) int {
	p := 1
	i := 1
	for i < len(s) && p > 0 {
		if s[i] == '(' {
			p++
		} else if s[i] == ')' {
			p--
		}
		i++
	}
	return i
}

func parse(r map[string]int) string {
	atoms := make([]string, 0, len(r))
	for a := range r {
		atoms = append(atoms, a)
	}
	sort.Strings(atoms)
	res := ""
	for _, a := range atoms {
		res += a
		if r[a] > 1 {
			res += strconv.Itoa(r[a])
		}
	}
	return res
}
