package Problem0736

import (
	"strconv"
	"strings"
)

func evaluate(expression string) int {
	expression = strings.Replace(expression, "(", "( ", -1)
	expression = strings.Replace(expression, ")", " )", -1)
	return helper(expression, nil)
}

func helper(exp string, s *scope) int {
	if exp[0] != '(' {
		num, err := strconv.Atoi(exp)
		if err == nil {
			return num
		}
		return s.get(exp)
	}

	// 删除最外层的 "( " 和 " )"
	exp = exp[2 : len(exp)-2]
	es := split(exp)
	switch es[0] {
	case "add":
		return helper(es[1], s) + helper(es[2], s)
	case "mult":
		return helper(es[1], s) * helper(es[2], s)
	default:
		s = newScope(s)
		var i int
		for i = 1; i < len(es)-2; i += 2 {
			s.add(es[i], helper(es[i+1], newScope(s)))
		}
		return helper(es[i], s)
	}

}

func split(exp string) []string {
	ss := strings.Split(exp, " ")
	countLeft := 0
	res := make([]string, 0, len(ss))
	for _, s := range ss {
		if countLeft == 0 {
			res = append(res, s)
		} else {
			res[len(res)-1] += " " + s
		}
		if s == "(" {
			countLeft++
		} else if s == ")" {
			countLeft--
		}
	}

	return res
}

type scope struct {
	parent *scope
	m      map[string]int
}

func newScope(parent *scope) *scope {
	return &scope{parent: parent, m: make(map[string]int, 8)}
}

func (s *scope) get(key string) int {
	if val, ok := s.m[key]; ok {
		return val
	}
	return s.parent.get(key)
}

func (s *scope) add(key string, val int) {
	s.m[key] = val
}
