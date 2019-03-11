package problem0948

import "sort"

// ref: https://leetcode.com/problems/bag-of-tokens/discuss/197696/C%2B%2BJavaPython-Greedy-%2B-Two-Pointers

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	scores := 0
	i, j := 0, len(tokens)-1

	for i <= j {
		if P < tokens[i] { // have no money to buy
			return scores
		}
		for i <= j && P >= tokens[i] { // buy the cheapest greedily
			P -= tokens[i]
			scores++
			i++
		}
		if i < j && scores > 0 { // sell the most expensive slowly
			P += tokens[j]
			scores--
			j--
		}
	}
	return scores
}
