package Problem0770

import (
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
	res := make(nums, 0, 128)
	for len(expression) > 0 {
		if expression[0] == '(' {
			i := indexOfCounterParentheses(expression)
			res = append(res, parse(expression[1:i], m))
			expression = expression[i+1:]
		}
	}
}

func indexOfCounterParentheses(expression string) int {
	for i := 1; i < len(expression); i++ {
		if expression[i] == ')' {
			return i
		}
	}
}

func getResult(numbers nums) []string {

}

type nums []num

type num struct {
	vars []string // 单个变量
	c    int      // 系数
}

func (n *num) update(m map[string]int) {
	for i, s := range n.vars {
		if v, ok := m[s]; ok {
			n.c *= v
			n.vars[i] = ""
		}
	}

}
