package Problem0466

type pair struct {
	fi int
	se int
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	dp := make(map[pair]pair)
	m1, m2, i, j := 0, 0, 0, 0
	for m1 < n1 {
		for s1[i] != s2[j] {
			i++
			if i == len(s1) {
				m1++
				i = 0
			}
			if m1 == n1 {
				return m2 / n2
			}
		}
		p := pair{i, j}
		if l, ok := dp[p]; ok {
			d := n1 - m1 - 1
			if d < 0 {
				d = 0
			}
			m2 += d / (m1 - l.fi) * (m2 - l.se)
			m1 += d / (m1 - l.fi) * (m1 - l.fi)
		}
		dp[p] = pair{m1, m2}

		i++
		j++
		if i == len(s1) {
			m1++
			i = 0
		}
		if j == len(s2) {
			m2++
			j = 0
		}
	}
	return m2 / n2
}
