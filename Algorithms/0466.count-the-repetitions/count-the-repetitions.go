package Problem0466

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	// count[k] == cnt 表示 k 个 s1 可以得到 cnt 个 s2
	count := make([]int, n1+1)
	// last[j] == k (k ∈ [1,2,...n1] ) 表示在第 k + 1 个 s1 中，首先需要匹配的是 s2[j]
	last := make([]int, len(s2))

	// j 是 s2 中下一个需要被匹配的字符在 s2 中的索引号
	// cnt 是 已经完整匹配了的 s2 的个数
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

		if last[j] == 0 {
			count[k] = cnt
			last[j] = k
		} else {
			start := last[j]
			p := k - start
			t := cnt - count[start]

			ans := (n1-start)/p*t + count[(n1-start)%p+start]
			return ans / n2
		}

	}

	return cnt / n2

}
