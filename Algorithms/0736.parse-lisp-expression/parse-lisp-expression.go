package Problem0736

import (
	"fmt"
	"strconv"
	"strings"
)

func evaluate(expression string) int {
	m := make(map[string]int, 8)
	return helper(expression, m)
}

func helper(exp string, m map[string]int) int {
	if exp[0] != '(' {
		num, err := strconv.Atoi(exp)
		if err == nil {
			return num
		}
		return m[exp]
	}

	// 删除最外层的 "()"
	exp = exp[1 : len(exp)-1]
	es := split(exp)
	switch es[0] {
	case "add":
		return helper(es[1], copy(m)) + helper(es[2], copy(m))
	case "mult":
		return helper(es[1], copy(m)) * helper(es[2], copy(m))
	default:
		var i int
		for i = 1; i < len(es)-2; i += 2 {
			m[es[i]] = helper(es[i+1], copy(m))
		}
		// TODO: 删除此处内容
		fmt.Println("\n=======\n", es[i])
		return helper(es[i], copy(m))
	}

}

func split(exp string) []string {
	exp = strings.Replace(exp, ")", " )", -1)
	// TODO: 删除此处内容
	fmt.Println(exp)
	ss := strings.Split(exp, " ")
	leftCount := 0
	res := make([]string, 0, len(ss))
	for _, s := range ss {
		if leftCount == 0 {
			res = append(res, s)
		} else {
			res[len(res)-1] += " " + s
		}
		// TODO: 删除此处内容
		fmt.Println(res[len(res)-1])

		if s[0] == '(' {
			leftCount++
		} else if s == ")" {
			leftCount--
		}
	}

	return res
}

func copy(m map[string]int) map[string]int {
	res := make(map[string]int, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}
