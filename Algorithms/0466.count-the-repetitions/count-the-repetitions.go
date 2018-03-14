package problem0466

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
			// 还没有出现重复
			// 更新各项记录
			count[k] = cnt
			last[j] = k
		} else { // 出现重复的模式了
			// 第 start 个和第 k 个 s1 都需要首先匹配 s[j]
			// p := k - start 的话，以此类推
			// 第 start + p, start + p * 2, start + p * 3,... 个 s1 都需要首先匹配 s[j]
			// 所以 s[start:k] 构成了一个重复块， 每个重复块中需要 p 个 s1
			start := last[j]
			p := k - start
			// 每个重复块可以匹配 t 个 s2
			// 此时的 cnt 相当于 count[k]
			t := cnt - count[start]

			// (n1-start)/p 是 n1 个 s1 可以组成的重复块的个数
			// (n1-start)/p*t 是所有完整的重复块可以匹配 s2 的个数
			// count[start+(n1-start)%p] 是剩下的，不能构成完整重复块的 s1 能够匹配的 s2 的个数
			ans := (n1-start)/p*t + count[start+(n1-start)%p]
			// 题目所求的 S2 的匹配数 = s2 的匹配数 / n2
			return ans / n2
		}
	}

	// 由于重复块不是一定存在
	return cnt / n2
}
