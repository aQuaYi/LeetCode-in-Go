package problem1096

import (
	"sort"
	"strings"
)

func braceExpansionII(exp string) []string {
	n := len(exp)

	bytes := make([]byte, 1, n*2)
	bytes[0] = exp[0]
	pre := exp[0]
	for i := 1; i < n; i++ {
		cur := exp[i]
		// add '*' symbol
		if (isLetter(pre) && cur == '{') ||
			(pre == '}' && cur == '{') ||
			(pre == '}' && isLetter(cur)) {
			bytes = append(bytes, '*')
		}
		// replace ',' with '+' symbol
		if cur == ',' {
			bytes = append(bytes, '+')
		} else {
			bytes = append(bytes, cur)
		}
		pre = cur
	}

	exp = string(bytes)

	return unique(doAdd(exp))
}

func isLetter(r byte) bool {
	return 'a' <= r && r <= 'z'
}

func split(exp string, topSymbol byte) []string {
	count := 0
	bytes := []byte(exp)
	for i, b := range bytes {
		switch b {
		case '{':
			count++
		case '}':
			count--
		case topSymbol:
			if count == 0 { // it's top now
				bytes[i] = '@'
			}
		}
	}
	exp = string(bytes)
	return strings.Split(exp, "@")
}

func doAdd(exp string) []string {
	exp = removeOuterBrace(exp)
	if !strings.ContainsRune(exp, '*') {
		exp = strings.Replace(exp, "{", "", -1)
		exp = strings.Replace(exp, "}", "", -1)
		return strings.Split(exp, "+")
	}
	exps := split(exp, '+')
	res := []string{}
	for _, e := range exps {
		res = add(res, doMultiply(e))
	}
	return res
}

func doMultiply(exp string) []string {
	exp = removeOuterBrace(exp)
	if !strings.ContainsRune(exp, '+') {
		exp = strings.Replace(exp, "{", "", -1)
		exp = strings.Replace(exp, "}", "", -1)
		exp = strings.Replace(exp, "*", "", -1)
		return []string{exp}
	}
	exps := split(exp, '*')
	res := []string{""}
	for _, e := range exps {
		res = multiply(res, doAdd(e))
	}
	return res
}

func removeOuterBrace(exp string) string {
	n := len(exp)
	if exp[0] != '{' ||
		exp[n-1] != '}' {
		return exp
	}
	count, i := 1, 1
	for ; i < n && count > 0; i++ {
		if exp[i] == '{' {
			count++
		} else if exp[i] == '}' {
			count--
		}
	}
	if i == n {
		return exp[1 : n-1]
	}
	return exp
}

func add(A, B []string) []string {
	return append(A, B...)
}

func multiply(A, B []string) []string {
	res := make([]string, 0, len(A)*len(B))
	for _, a := range A {
		for _, b := range B {
			res = append(res, a+b)
		}
	}
	return res
}

func unique(A []string) []string {
	sort.Strings(A)
	i, j, n := 0, 0, len(A)
	for j < n {
		if A[i] != A[j] {
			i++
			A[i] = A[j]
		}
		j++
	}
	return A[:i+1]
}
