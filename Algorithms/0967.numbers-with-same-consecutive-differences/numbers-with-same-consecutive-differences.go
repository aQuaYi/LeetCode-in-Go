package problem0967

func numsSameConsecDiff(N int, K int) []int {
	res := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if N == 1 {
		return res
	}

	res = make([]int, 0, 20)
	for i := 1; i < 10; i++ {
		if t, ok := makeNumber(i, N, -K); ok {
			res = append(res, t)
		}
		if K == 0 {
			continue
		}
		if t, ok := makeNumber(i, N, K); ok {
			res = append(res, t)
		}
	}

	return res
}

func makeNumber(x, N, K int) (int, bool) {
	if x+K < 0 || 9 < x+K {
		return 0, false
	}
	for i := 1; i < N; i++ {
		x = x*10 + x%10 + K
		K = -K
	}
	return x, true
}
