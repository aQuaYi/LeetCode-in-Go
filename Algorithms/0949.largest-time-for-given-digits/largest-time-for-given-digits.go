package problem0949

import "fmt"

var permutations = [][]int{
	{0, 1, 2, 3},
	{0, 1, 3, 2},
	{0, 2, 1, 3},
	{0, 2, 3, 1},
	{0, 3, 1, 2},
	{0, 3, 2, 1},
	{1, 0, 2, 3},
	{1, 0, 3, 2},
	{1, 2, 0, 3},
	{1, 2, 3, 0},
	{1, 3, 0, 2},
	{1, 3, 2, 0},
	{2, 0, 1, 3},
	{2, 0, 3, 1},
	{2, 1, 0, 3},
	{2, 1, 3, 0},
	{2, 3, 1, 0},
	{2, 3, 0, 1},
	{3, 0, 2, 1},
	{3, 0, 1, 2},
	{3, 1, 0, 2},
	{3, 1, 2, 0},
	{3, 2, 0, 1},
	{3, 2, 1, 0},
}

func largestTimeFromDigits(A []int) string {
	h, m := -1, -1
	res := ""
	for _, p := range permutations {
		a, b, c, d := A[p[0]], A[p[1]], A[p[2]], A[p[3]]
		ab, cd := a*10+b, c*10+d
		if ab > 23 ||
			cd > 59 ||
			ab*60+cd <= h*60+m {
			continue
		}
		h, m = ab, cd
		res = fmt.Sprintf("%02d:%02d", h, m)
	}
	return res
}
