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
			maxScore = score
			// 因为 tokens 是升序排列
			// 所以，卖了 tokens[j] 后，肯定应该买得起 tokens[i]
			// 除非 i==j，会导致 score-- 后， score 比真正的结果小 1
			// 于是，每次 score++ 后，score 肯定是最大值。
			// 使用 maxScore 记录一下，避免错误的结果。
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
