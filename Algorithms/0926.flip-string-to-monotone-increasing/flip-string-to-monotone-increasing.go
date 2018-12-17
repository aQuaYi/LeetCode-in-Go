package problem0926

func minFlipsMonoIncr(S string) int {
	size := len(S)

	count := make([]int, size+1)
	for i := 0; i < size; i++ {
		count[i+1] = count[i]
		if S[i] == '1' {
			count[i+1]++
		}
	}

	res := size
	for i := 0; i <= size; i++ {
		// for S[:i] is all 0
		// need flip 1 to 0 in S[:i],
		//      there are count[i] 1s in S[i:]
		// for S[i:] is all 1
		// need flip 0 to 1 in S[i:],
		//      length of S[i:] is size-i,
		//      count of 1s in S[i:] is count[size]-count[i],
		//      there are (size-i)-(count[size]-count[i]) 0s in S[i:]
		tmp := count[i] + ((size - i) - (count[size] - count[i]))
		res = min(res, tmp)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
