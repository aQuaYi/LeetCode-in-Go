package problem1072

import "strings"

func maxEqualRowsAfterFlips(matrix [][]int) int {
	count := make(map[string]int, 300)
	xor := make(map[string]string, 300)
	for _, m := range matrix {
		a, b := convert(m)
		count[a]++
		xor[a] = b
	}

	res := 0

	for a, c := range count {
		res = max(res, c+count[xor[a]])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func convert(nums []int) (string, string) {
	var x, xor strings.Builder
	for _, n := range nums {
		x.WriteByte(byte(n + '0'))
		xor.WriteByte(byte(1 - n + '0'))
	}
	return x.String(), xor.String()
}
