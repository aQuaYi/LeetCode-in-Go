package Problem0763

func partitionLabels(S string) []int {
	maxIndex := [26]int{}
	for i, b := range S {
		maxIndex[b-'a'] = i
	}

	begin := 0
	end := maxIndex[S[begin]-'a']
	res := make([]int, 0, len(S))

	for i, b := range S {
		if end == maxIndex[b-'a'] {
			res = append(res, end-begin+1)
			begin = i + 1
			if begin < len(S) {
				end = maxIndex[S[begin]-'a']
			} else {
				res = append(res, len(S)-i-1)
			}
		}
	}

	return res
}
