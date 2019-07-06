package problem1044

import "math"

// ref: https://leetcode.com/problems/longest-duplicate-substring/discuss/290871/Python-Binary-Search
func longestDupSubstring(S string) string {
	A := numberify(S)
	mod := 1<<63 - 1

	test := func(L int) int {
		p := int(math.Mod(math.Pow(26, float64(L)), float64(mod)))
		cur := 0
		for i := 0; i < L; i++ {
			cur = (cur*26 + A[i]) % mod
		}
		seen := make(map[int]bool, len(S)-L)
		seen[cur] = true
		for i := L; i < len(S); i++ {
			cur = (cur*26 + A[i] - A[i-L]*p) % mod
			if seen[cur] {
				return i - L + 1
			}
			seen[cur] = true
		}
		return 0
	}

	res, lo, hi := 0, 0, len(S)
	for lo < hi {
		mi := (lo + hi + 1) / 2
		pos := test(mi)
		if pos > 0 {
			lo = mi
			res = pos
		} else {
			hi = mi - 1
		}
	}

	return S[res : res+lo]
}

func numberify(S string) []int {
	res := make([]int, len(S))
	for i, r := range S {
		res[i] = int(r - 'a')
	}
	return res
}
