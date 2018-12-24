package problem0940

const mod = 1e9 + 7

func distinctSubseqII(S string) int {
	tail := [26]int{}

	sum := func() int {
		count := 0
		for i := 0; i < 26; i++ {
			count += tail[i]
		}
		return count % mod
	}

	for _, r := range S {
		i := int(r - 'a')
		tail[i] = sum() + 1
	}

	return sum()
}
