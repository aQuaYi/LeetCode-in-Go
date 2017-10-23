package Problem0354

import "sort"

func maxEnvelopes(es [][]int) int {
	size := len(es)
	if size == 0 {
		return 0
	}

	e := envelopes(es)
	sort.Sort(e)

	// dp[i] 表示 envelopes[i] 能放入的最大 doll
	dp := make([]int, size)

	var i, j int
	for i = size - 2; 0 <= i; i-- {
		for j = size - 1; i < j; j-- {
			if isFit(e[i], e[j]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return dp[0] + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// return true 当 ei 可以放入 ej 的时候
func isFit(ei, ej []int) bool {
	return ei[0] < ej[0] && ei[1] < ej[1]
}

type envelopes [][]int

func (e envelopes) Len() int {
	return len(e)
}

func (e envelopes) Less(i, j int) bool {
	if e[i][0] == e[j][0] {
		return e[i][1] < e[j][1]
	}
	return e[i][0] < e[j][0]
}

func (e envelopes) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
