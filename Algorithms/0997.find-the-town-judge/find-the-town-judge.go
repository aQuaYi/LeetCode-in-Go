package problem0997

func findJudge(N int, trust [][]int) int {
	count := [1001]int{}
	trustSomebody := [1001]bool{}
	for _, t := range trust {
		count[t[1]]++
		trustSomebody[t[0]] = true
	}

	for i := 1; i <= N; i++ {
		if count[i] == N-1 && !trustSomebody[i] {
			return i
		}
	}

	return -1
}
