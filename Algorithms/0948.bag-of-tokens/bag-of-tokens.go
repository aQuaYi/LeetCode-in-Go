package problem0948

import "sort"

// ref: https://leetcode.com/problems/bag-of-tokens/discuss/197696/C%2B%2BJavaPython-Greedy-%2B-Two-Pointers

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	res, scores := 0, 0
	i, j := 0, len(tokens)-1

	for i <= j {
		if P >= tokens[i] {
			P -= tokens[i]
			scores++
			res = max(res, scores)
			i++
		} else if scores > 0 {
			P += tokens[j]
			scores--
			j--
		} else {
			break
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
