package problem0967

func numsSameConsecDiff(N int, K int) []int {
	res := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := 1; n < N; n++ {
		tmp := make([]int, 0, len(res)*2)
		for _, v := range res {
			if v == 0 {
				continue
			}
			r := v % 10
			if r-K >= 0 {
				tmp = append(tmp, v*10+r-K)
			}
			if r+K <= 9 {
				tmp = append(tmp, v*10+r+K)
			}
		}
		res = tmp
	}
	return res
}
