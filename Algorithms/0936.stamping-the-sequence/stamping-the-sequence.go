package problem0936

import "strings"

func movesToStamp(stamp string, target string) []int {

	n, m, t, s := len(target), len(stamp), []byte(target), []byte(stamp)
	res := make([]int, 0, n)

	check := func(i int) bool {
		changed := false
		for j := 0; j < m; j++ {
			if t[i+j] == '?' {
				continue
			}
			if t[i+j] != s[j] {
				return false
			}
			changed = true
		}

		if changed {
			for k := i; k < i+m; k++ {
				t[k] = '?'
			}
			res = append(res, i)
		}
		return changed
	}

	changed := true
	for changed {
		changed = false
		for i := 0; i < n-m+1; i++ {
			changed = changed || check(i)
		}
	}

	reverse(res)

	if string(t) == strings.Repeat("?", n) {
		return res
	}
	return nil
}

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}
