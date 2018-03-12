package Problem0770

import (
	"sort"
	"strings"
)

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	m := make(map[string]int, len(evalvars))
	for i := range evalvars {
		m[evalvars[i]] = evalints[i]
	}

	expression = strings.Replace(expression, " ", "", -1)

	numbers := parse(expression, m)

	return getResult(numbers)
}

func parse(expression string, m map[string]int) nums {
	res := make(nums, 1, 128)
	res[0] = num{c: 0}
	opt = '+'
	for len(expression) > 0 {
		var next nums
		var i int
		if expression[0] == '(' {
			i = indexOfCounterParentheses(expression)
			next = parse(expression[1:i], m)
		} else {
			i := indexOfNextOpt(expression)
			n := expression[:i]
			v, ok := m[n]
			if ok {
				next = nums{num{c: v}}
			} else {
				next = nums{num{vars: []string{n}}}
			}
		}

	}
}

func operate(opt byte, a, b nums) nums {
	var res nums
	switch opt {
	case '+':
		res = append(a, b...)
	case '-':
		res = minus(a, b)
	case '*':
		res = mult(a, b)
	}
}

func minus(a, b nums) nums {
	for i := range b {
		b[i].c *= -1
	}
	return append(a, b...)
}

func mult(a, b nums) nums {
	res := make(nums, 0, len(a)*len(b))
	for i := range a {
		for j := range b {
			vars := append(a[i].vars, b[i].vars...)
			res = append(res, num{vars: vars, c: a[i].c * b[j].c})
		}
	}

	return res
}

func indexOfCounterParentheses(expression string) int {
	for i := 1; i < len(expression); i++ {
		if expression[i] == ')' {
			return i
		}
	}
}

func indexOfNextOpt(expression string) int {
	var i int
	for i = 1; i < len(expression); i++ {
		if expression[i] == '+' ||
			expression[i] == '-' ||
			expression[i] == '*' {
			break
		}
	}
	return i
}

func getResult(numbers nums) []string {

	return nil
}

type nums []num

type num struct {
	vars []string // 变量
	key  string
	c    int // 系数
}

func update(ns nums) nums {
	for i := range ns {
		sort.Strings(ns[i].vars)
		ns[i].key = strings.Join(ns[i].vars, "*")
	}

	sort.Slice(ns, func(i int, j int) bool {
		return ns[i].key < ns[j].key
	})

	res := make(nums, 1, len(ns))
	res[0] = ns[0]
	i, j := 0, 1
	for j < len(ns) {
		if res[i].key != ns[j].key {
			res = append(res, ns[j])
			i++
			j++
			continue
		}

		res[i].c += ns[j].c
		j++
	}

	return res
}
