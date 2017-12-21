package Problem0673

func findNumberOfLIS(a []int) int {
	res := 0
	maxLen := 0

	dfs(a, -1<<63, 0, 0, &maxLen, &res)

	return res
}

func dfs(a []int, Min, index, length int, maxLen, res *int) {
	if *maxLen < length {
		*maxLen = length
		*res = 1
	} else if *maxLen == length && length > 0 {
		*res++
	}

	Max := 1<<63 - 1

	for i := index; i < len(a); i++ {
		if Min < a[i] && a[i] <= Max {
			Max = a[i]
			dfs(a, a[i], i+1, length+1, maxLen, res)
		}
	}
}
