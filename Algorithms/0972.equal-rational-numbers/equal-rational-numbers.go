package problem0972

import (
	"strconv"
	"strings"
)

func isRationalEqual(S string, T string) bool {
	is, ns, rs := parse(S)
	it, nt, rt := parse(T)

	if rs == "" && rt == "" {
		return S == T
	}

	if rs == "" || rt == "" {
		return false
	}

	ns, rs = simplify(ns, rs)
	nt, rt = simplify(nt, rt)

	s, ns, rs := convert(is, ns, rs)
	t, nt, rt := convert(it, nt, rt)

	return s == t && ns == nt && rs == rt
}

func parse(s string) (string, string, string) {
	dot := strings.Index(s, ".")
	if dot == -1 {
		return s, "", ""
	}

	integer, fraction := s[:dot], s[dot+1:]

	l := strings.Index(fraction, "(")
	if l == -1 {
		return integer, fraction, ""
	}

	nonRepeat := fraction[:l]
	repeat := fraction[l+1 : len(fraction)-1]

	return integer, nonRepeat, repeat
}

func simplify(nonRepeat, repeat string) (string, string) {
	for repeat[:len(repeat)/2] == repeat[len(repeat)/2:] {
		repeat = repeat[:len(repeat)/2]
	}

	for strings.HasSuffix(nonRepeat, repeat) {
		nonRepeat = nonRepeat[:len(nonRepeat)-len(repeat)]
	}

	return nonRepeat, repeat
}

func convert(integer, nonRepeat, repeat string) (int, string, string) {
	i, _ := strconv.Atoi(integer)
	if nonRepeat == "" && repeat == "9" {
		i++
		repeat = ""
	}
	return i, nonRepeat, repeat
}
