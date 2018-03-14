package problem0721

import (
	"sort"
)

func accountsMerge(acts [][]string) [][]string {
	n := len(acts)

	owner := make(map[string]string, n)
	parent := make(map[string]string, n*2)

	for _, a := range acts {
		for i := 1; i < len(a); i++ {
			parent[a[i]] = a[i]
			owner[a[i]] = a[0]
		}
	}

	for _, a := range acts {
		r := root(a[1], parent)
		for i := 2; i < len(a); i++ {
			p := root(a[i], parent)
			parent[p] = r
		}
	}

	union := make(map[string][]string, n)
	for email, p := range parent {
		r := root(p, parent)
		union[r] = append(union[r], email)
	}

	res := make([][]string, 0, len(union))
	for p, emails := range union {
		t := make([]string, len(emails)+1)
		t[0] = owner[p]
		if len(emails) > 1 {
			sort.Strings(emails)
		}
		copy(t[1:], emails)
		res = append(res, t)
	}

	return res
}

func root(e string, parent map[string]string) string {
	if parent[e] == e {
		return e
	}
	return root(parent[e], parent)
}
