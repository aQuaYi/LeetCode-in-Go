package problem1072

import "strings"

func maxEqualRowsAfterFlips(A [][]int) int {
	count := make(map[string]int, 300)
	var sb strings.Builder
	for _, r := range A {
		r0 := r[0]
		for _, x := range r {
			sb.WriteByte(byte(x ^ r0 + '0'))
		}
		count[sb.String()]++
		sb.Reset()
	}

	res := 0

	for _, c := range count {
		res = max(res, c)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
