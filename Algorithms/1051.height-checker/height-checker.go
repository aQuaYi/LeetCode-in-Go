package problem1051

import "sort"

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)
	res := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sorted[i] {
			res++
		}
	}
	return res
}
