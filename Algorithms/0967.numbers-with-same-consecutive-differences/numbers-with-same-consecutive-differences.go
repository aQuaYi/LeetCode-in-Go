package problem0967

func numsSameConsecDiff(N int, K int) []int {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if N == 1 {
		// when N is 1, whatever K is, the answer is nums
		return nums
	}

	res := nums[1:] // when N>1, delete 0
	// bfs
	for n := 1; n < N; n++ {
		tmp := make([]int, 0, len(res)*2)
		for _, v := range res {
			r := v % 10
			if r-K >= 0 {
				tmp = append(tmp, v*10+r-K)
			}
			if K > 0 && r+K <= 9 {
				tmp = append(tmp, v*10+r+K)
			}
		}
		res = tmp
	}

	return res
}
