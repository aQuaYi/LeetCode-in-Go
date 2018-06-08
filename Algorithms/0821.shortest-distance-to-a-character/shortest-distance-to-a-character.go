package problem0821

func shortestToChar(S string, C byte) []int {
	n := len(S)

	res := make([]int, n)
	for i := range res {
		res[i] = n
	}

	left, right := -n, 2*n

	for i := 0; i < n; i++ {
		j := n - i - 1
		if S[i] == C {
			left = i
		}
		if S[j] == C {
			right = j
		}
		res[i] = min(res[i], dist(i, left))
		res[j] = min(res[j], dist(j, right))
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dist(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
