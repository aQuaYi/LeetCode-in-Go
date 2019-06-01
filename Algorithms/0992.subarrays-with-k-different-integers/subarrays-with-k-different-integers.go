package problem0992

// ref: https://leetcode.com/problems/subarrays-with-k-different-integers/discuss/234482/JavaC%2B%2BPython-Sliding-Window-with-Video
func subarraysWithKDistinct(A []int, K int) int {
	return atMostK(A, K) - atMostK(A, K-1)
}

func atMostK(A []int, K int) int {
	count := [20001]int{}
	res := 0
	for l, r := 0, 0; r < len(A); r++ {
		if count[A[r]] == 0 {
			K--
		}
		count[A[r]]++
		for K < 0 {
			count[A[l]]--
			if count[A[l]] == 0 {
				K++
			}
			l++
		}
		// the number of distinct elements in A[i:r+1] (l<=i<=r)
		// is NOT more than K
		res += r + 1 - l
	}
	return res
}
