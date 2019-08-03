package problem1106

import "strings"

func parseBoolExpr(exp string) bool {
	if exp == "t" {
		return true
	}
	if exp == "f" {
		return false
	}
	n := len(exp)
	symbol, exp := exp[0], exp[2:n-1]
	subs := split(exp)
	switch symbol {
	case '&':
		return and(subs)
	case '|':
		return or(subs)
	default:
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

func or(exps []string) bool {
	for _, e := range exps {
		if parseBoolExpr(e) {
			return true
		}
	}
	return false
}

func and(exps []string) bool {
	for _, e := range exps {
		if !parseBoolExpr(e) {
			return false
		}
	}
	return true
}

func not(exp string) bool {
	return !parseBoolExpr(exp)
}
