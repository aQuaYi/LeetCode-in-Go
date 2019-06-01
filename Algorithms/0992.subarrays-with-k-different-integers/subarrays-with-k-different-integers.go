package problem0992

func subarraysWithKDistinct(A []int, K int) int {
	counts, k := [20001]int{}, 0

	size := len(A)
	l, r := 0, -1
	res := 0
	for r+1 < size {
		for k < K && r+1 < size {
			r++
			if counts[A[r]] == 0 {
				k++
			}
			counts[A[r]]++
		}

		for k == K {
			res++
			counts[A[l]]--
			if counts[A[l]] == 0 {
				k--
			}
			l++
		}
	}

	return res
}
