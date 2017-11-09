package Problem0466

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {

	repeatCount := make([]int, len(s2)+1)
	visited := make([]int, len(s2)+1)

	j, cnt := 0, 0

	for k := 1; k <= n1; k++ {
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[j] {
				j++
				if j == len(s2) {
					j = 0
					cnt++
				}
			}
		}

		if visited[j] == 0 {
			repeatCount[k] = cnt
			visited[j] = k
		} else {
			start := visited[j]
			p := k - start
			t := cnt - repeatCount[start]

			ans := (n1-start)/p*t + repeatCount[(n1-start)%p+start]
			return ans / n2
		}

	}

	return cnt / n2

}
