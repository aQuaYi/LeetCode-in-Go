package problem0972

import (
	"strconv"
	"strings"
)

func isRationalEqual(S string, T string) bool {
	S = normalize(S)
	sn, sr := parse2(S)
	sn, sr = simplify2(sn, sr)
	s, sr := convert2(sn, sr)

	T = normalize(T)
	tn, tr := parse2(T)
	tn, tr = simplify2(tn, tr)
	t, tr := convert2(tn, tr)

	return s == t && sr == tr
}

// 1.2(3) -> 0.00012(3) -> 00012(3)
func normalize(s string) string {
	if !strings.Contains(s, ".") {
		s += "."
	}
	if strings.HasSuffix(s, ".") {
		s += "0"
	}
	dot := strings.Index(s, ".")
	if dot < 4 {
		s = strings.Repeat("0", 4-dot) + s
	}
	return strings.Replace("1"+s, ".", "", 1)
}

func parse2(s string) (string, string) {
	if !strings.Contains(s, "(") {
		return s, ""
	}

	i := strings.Index(s, "(")

	nonRepeat := s[:i]
	repeat := s[i+1 : len(s)-1]

	if repeat == "0" {
		repeat = ""
	}

	return nonRepeat, repeat
}

func parse(s string) (string, string, string) {
	if strings.HasSuffix(s, ".") {
		s = s[:len(s)-1]
	}

	dot := strings.Index(s, ".")
	if dot == -1 {
		return s, "", ""
	}

	integer, fraction := s[:dot], s[dot+1:]

	l := strings.Index(fraction, "(")
	if l == -1 {
		if fraction == "0" {
			return integer, "", ""
		}
		return integer, fraction, ""
	}

	nonRepeat := fraction[:l]
	repeat := fraction[l+1 : len(fraction)-1]

	if repeat == "0" {
		repeat = ""
	}

	if repeat == "" && nonRepeat == "0" {
		nonRepeat = ""
	}

	return integer, nonRepeat, repeat
}

func simplify2(nonRepeat, repeat string) (string, string) {
	if repeat == "" {
		return nonRepeat, repeat
	}

	if repeat == strings.Repeat(repeat[:1], len(repeat)) {
		repeat = repeat[:1]
	}

	for repeat[:len(repeat)/2] == repeat[len(repeat)/2:] {
		repeat = repeat[:len(repeat)/2]
	}

	for strings.HasSuffix(nonRepeat, repeat) {
		nonRepeat = nonRepeat[:len(nonRepeat)-len(repeat)]
	}

	for i := 1; i < len(repeat); i++ {
		if strings.HasSuffix(nonRepeat, repeat[i:]) {
			repeat = repeat[i:] + repeat[:i]
			nonRepeat = nonRepeat[:len(nonRepeat)-len(repeat)+i]
			break
		}
	}

	return nonRepeat, repeat
}

func simplify(nonRepeat, repeat string) (string, string) {
	if repeat == "" {
		return nonRepeat, repeat
	}

	if repeat == strings.Repeat(repeat[:1], len(repeat)) {
		repeat = repeat[:1]
	}

	for repeat[:len(repeat)/2] == repeat[len(repeat)/2:] {
		repeat = repeat[:len(repeat)/2]
	}

	for strings.HasSuffix(nonRepeat, repeat) {
		nonRepeat = nonRepeat[:len(nonRepeat)-len(repeat)]
	}

	for i := 1; i < len(repeat); i++ {
		if strings.HasSuffix(nonRepeat, repeat[i:]) {
			repeat = repeat[i:] + repeat[:i]
			nonRepeat = nonRepeat[:len(nonRepeat)-len(repeat)+i]
			break
		}
	}

	return nonRepeat, repeat
}

func convert2(nonRepeat, repeat string) (int, string) {
	i, _ := strconv.Atoi(nonRepeat)
	if repeat == "9" {
		i++
		repeat = ""
	}
	return i, repeat
}

func convert(integer, nonRepeat, repeat string) (int, string) {
	i, _ := strconv.Atoi(integer)
	for j := len(nonRepeat); j > 0; j-- {
		i *= 10
	}
	if nonRepeat != "" {
		n, _ := strconv.Atoi(nonRepeat)
		i += n
	}
	if repeat == "9" {
		i++
		repeat = ""
	}
	return i, repeat
}
