package problem0955

func minDeletionSize(A []string) int {
	m, n := len(A), len(A[0])
	res := 0

	for j := 0; j < n; j++ {
		i := 1
		t := 0
		for ; i < m; i++ {
			p, c := A[i-1][j], A[i][j]
			if p > c {
				break
			} else if p == c {
				t = i + 1
			}
		}
		if i == m {
			if t == 0 {
				return res
			}
			m = t
		} else {
			res++
		}
	}
	return n
}

func check(s1, s2 string, marks []int) {
	for i, m := range marks {
		if m == 1 { // deleted column
			continue
		}
		switch {
		case s1[i] < s2[i]:
			return
		case s1[i] > s2[i]:
			marks[i] = 1

		}
	}
}

func sum(marks []int) int {
	sum := 0
	for _, m := range marks {
		sum += m
	}
	return sum
}
