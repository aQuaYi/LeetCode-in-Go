package problem0936

import "strings"

func movesToStamp(stamp string, target string) []int {
	t, s := []byte(target), []byte(stamp)
	tSize, sSize := len(t), len(s)

	res := make([]int, 0, tSize)

	isStamped := func(i int) bool {
		changed := false
		for j := 0; j < sSize; j++ {
			if t[i+j] == '?' {
				continue
			}
			if t[i+j] != s[j] {
				return false
			}
			changed = true
		}
		if changed {
			for k := i; k < i+sSize; k++ {
				t[k] = '?'
			}
			res = append(res, i)
		}
		return changed
	}

	changed := true
	for changed {
		changed = false
		for i := 0; i < tSize-sSize+1; i++ {
			changed = changed || isStamped(i)
		}
	}

	reverse(res)

	if string(t) == strings.Repeat("?", tSize) {
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
