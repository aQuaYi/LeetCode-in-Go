package problem0940

// ref: https://www.jianshu.com/p/02501f516437

const mod = 1e9 + 7

func distinctSubseqII(S string) int {
	tail := make([]int, 26)

	total := func() int {
		count := 0
		for i := 0; i < 26; i++ {
			count += tail[i]
		}
		return count % mod
	}

	for _, r := range S {
		tail[r-'a'] = total() + 1
	}

	return total()
}
