package problem0763

func partitionLabels(S string) []int {
	maxIndex := [26]int{}
	for i, b := range S {
		maxIndex[b-'a'] = i
	}

	begin := 0
	end := maxIndex[S[begin]-'a']
	res := make([]int, 0, len(S))

	for i, b := range S {
		if i < end {
			// 在 S[:i+1] 和 S[i:] 中存在相同的字母 S[end]
			// 所以此时不能分隔，仅更新 end
			end = max(end, maxIndex[b-'a'])
			continue
		}

		// 此时 S[begin:i+1] 中的所有字母都不会出现在其他片段中
		// 可以进行分隔
		res = append(res, i-begin+1)
		begin = i + 1 // 从 i+1 处作为新片段的起始点
		if begin < len(S) {
			// 及时更新 end
			end = maxIndex[S[begin]-'a']
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
