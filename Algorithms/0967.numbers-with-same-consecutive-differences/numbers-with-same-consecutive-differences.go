package problem0967

func numsSameConsecDiff(N int, K int) []int {
	queue := prepare(K)
	for n := 1; n < N; n++ {
		tmp := make([]int, 0, len(queue)*2)
		for _, v := range queue {
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
		queue = tmp
	}
	return queue
}

func prepare(k int) []int {
	res := make([]int, 0, 32)
	for i := 0; i < 10; i++ {
		if i+k > 9 && i-k < 0 {
			continue
		}
		res = append(res, i)
	}
	return res
}
