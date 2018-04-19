package problem0811

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int, len(cpdomains))

	for _, domin := range cpdomains {
		d, n := parse(domin)
		isNew := true
		for isNew {
			m[d] += n
			d, isNew = cut(d)
		}
	}

	return getResult(m)
}

func cut(s string) (string, bool) {
	idx := strings.Index(s, ".")
	if idx == -1 {
		return "", false
	}
	return s[idx+1:], true
}

func parse(s string) (string, int) {
	ss := strings.Split(s, " ")
	n, _ := strconv.Atoi(ss[0])
	return ss[1], n
}

func getResult(m map[string]int) []string {
	res := make([]string, 0, len(m))
	for k, v := range m {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}
