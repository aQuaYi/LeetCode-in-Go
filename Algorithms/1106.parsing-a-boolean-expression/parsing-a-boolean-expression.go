package problem1106

import "strings"

func parseBoolExpr(exp string) bool {
	if exp == "t" || exp == "f" {
		return exp == "t"
	}
	n := len(exp)
	op, exp := exp[0], exp[2:n-1]
	switch op {
	case '&':
		return and(split(exp))
	case '|':
		return or(split(exp))
	default: // !
		return not(exp)
	}
}

func split(exp string) []string {
	count := 0
	isUnbracketed := func() bool {
		return count == 0
	}
	bytes := []byte(exp)
	for i, b := range bytes {
		switch b {
		case '(':
			count++
		case ')':
			count--
		case ',':
			if isUnbracketed() {
				bytes[i] = '@'
			}
		}
	}
	exp = string(bytes)
	return strings.Split(exp, "@")
}

func and(exps []string) bool {
	for _, e := range exps {
		if !parseBoolExpr(e) {
			return false
		}
	}
	return true
}

func or(exps []string) bool {
	for _, e := range exps {
		if parseBoolExpr(e) {
			return true
		}
	}
	return false
}

func not(exp string) bool {
	return !parseBoolExpr(exp)
}
