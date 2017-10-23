package Problem0363

import (
	"sort"
)

func maxSumSubmatrix(mat [][]int, k int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}

	m, n := len(mat), len(mat[0])
	// 本算法的复杂度是 O(M*M*N*logN)
	// 让 N = max(m,n) 可以加快程序
	M, N := min(m, n), max(m, n)

	ans := 0

	var findMaxArea func([]int, int, int) int
	findMaxArea = func(sums []int, beg, end int) int {
		if beg+1 >= end {
			return -1 << 63
		}

		mid := (beg + end) >> 1

		res := max(findMaxArea(sums, beg, mid), findMaxArea(sums, mid, end))

		i := mid
		for _, l := range sums[beg:mid] {
			for i < len(sums) && sums[i]-l <= k {
				res = max(res, sums[i]-l)
				i++
			}
		}
		sort.Ints(sums[beg:end])
		return res
	}

	var i1, i2, j int
	for i1 = 0; i1 < M; i1++ {
		tmp := make([]int, N)
		for i2 = i1; i2 < M; i2++ {
			sums := []int{0}
			low := 0
			maxArea := -1 << 63
			for j = 0; j < N; j++ {
				if m < n {
					tmp[j] += mat[i2][j]
				} else {
					tmp[j] += mat[j][i2]
				}

				sums = append(sums, sums[len(sums)-1]+tmp[j])

				maxArea = max(maxArea, sums[len(sums)-1]-low)

				low = min(low, sums[len(sums)-1])

			}
			if maxArea <= ans {
				continue
			}
			if maxArea == k {
				return k
			}
			if maxArea > k {
				maxArea = findMaxArea(sums, 0, N+1)
			}
			ans = max(ans, maxArea)
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
