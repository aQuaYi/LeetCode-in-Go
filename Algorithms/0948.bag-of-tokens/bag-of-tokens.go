package problem0948

import "sort"

// ref: https://leetcode.com/problems/bag-of-tokens/discuss/197696/C%2B%2BJavaPython-Greedy-%2B-Two-Pointers

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	points := 0
	i, j := 0, len(tokens)-1
	for i <= j {
		if P < tokens[i] {
			return points
		}
		for i <= j && P >= tokens[i] {
			P -= tokens[i]
			points++
			i++
		}
		if i < j && points > 0 {
			P += tokens[j]
			points--
			j--
		}
	}
	return points
}
