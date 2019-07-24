package problem1096

import (
	"sort"
	"strings"
)

var pres = []string{"a{", "b{", "c{", "d{", "e{", "f{", "g{", "h{", "i{", "j{", "k{", "l{", "m{", "n{", "o{", "p{", "q{", "r{", "s{", "t{", "u{", "v{", "w{", "x{", "y{", "z{", "}{", "}a", "}b", "}c", "}d", "}e", "}f", "}g", "}h", "}i", "}j", "}k", "}l", "}m", "}n", "}o", "}p", "}q", "}r", "}s", "}t", "}u", "}v", "}w", "}x", "}y", "}z"}
var nows = []string{"a*{", "b*{", "c*{", "d*{", "e*{", "f*{", "g*{", "h*{", "i*{", "j*{", "k*{", "l*{", "m*{", "n*{", "o*{", "p*{", "q*{", "r*{", "s*{", "t*{", "u*{", "v*{", "w*{", "x*{", "y*{", "z*{", "}*{", "}*a", "}*b", "}*c", "}*d", "}*e", "}*f", "}*g", "}*h", "}*i", "}*j", "}*k", "}*l", "}*m", "}*n", "}*o", "}*p", "}*q", "}*r", "}*s", "}*t", "}*u", "}*v", "}*w", "}*x", "}*y", "}*z"}

func braceExpansionII(exp string) []string {
	for i, p := range pres {
		exp = strings.Replace(exp, p, nows[i], -1)
	}
	exp = strings.Replace(exp, ",", "+", -1)

	return unique(doAdd(exp))
}

func split(exp string, topSymbol byte) []string {
	exp = outBrace(exp)
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
	if !strings.ContainsRune(exp, '{') {
		return strings.Split(exp, "+")
	}
	//
	strs := split(exp, '+')
	res := []string{}
	for _, s := range strs {
		res = add(res, doMultiply(s))
	}
	return res
}

func doMultiply(exp string) []string {
	strs := split(exp, '*')
	res := []string{""}
	for _, s := range strs {
		res = multiply(res, doAdd(s))
	}
	return res
}

func outBrace(exp string) string {
	n := len(exp)
	if exp[0] != '{' ||
		exp[n-1] != '}' {
		return exp
	}
	count := 1
	i := 1
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
