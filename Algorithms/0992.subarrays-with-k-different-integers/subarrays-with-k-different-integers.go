package problem0992

func subarraysWithKDistinct(A []int, K int) int {
	counts, k := [20001]int{}, 0

	size := len(A)
	l, r := 0, -1
	res := 0
	for r+1 < size {
		c := 0
		for k <= K && r+1 < size {
			if k == K && counts[A[r+1]] == 0 {
				break
			}
			r++
			if counts[A[r]] == 0 {
				k++
			}
			counts[A[r]]++
			if k == K {
				c++
			}
		}

		if c > 1 {
			c--
		}

		for k == K {
			res += c
			counts[A[l]]--
			if counts[A[l]] == 0 {
				k--
			}
			l++
		}
	}

	return res
}
