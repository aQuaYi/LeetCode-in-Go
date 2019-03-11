package problem0948

import "sort"

// ref: https://leetcode.com/problems/bag-of-tokens/discuss/197696/C%2B%2BJavaPython-Greedy-%2B-Two-Pointers

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	maxScore, score := 0, 0
	i, j := 0, len(tokens)-1

	for i <= j {
		if P >= tokens[i] {
			P -= tokens[i]
			score++
			maxScore = max(maxScore, score) // record the largest score
			i++
		} else if score > 0 {
			P += tokens[j]
			score--
			j--
		} else {
			break
		}
	}

	return maxScore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
