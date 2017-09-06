package Problem0072

func minDistance(word1 string, word2 string) int {
	// 转换是可逆的
	// 互换可以只实现一套逻辑
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}

	targetSize := len(word2)
	if targetSize == 0 {
		return 0
	}

	origin := []byte(word1)
	target := []byte(word2)

	originSize := len(word1)
	if originSize == targetSize {
		return countDiff(origin, target)
	}

	match := 0
	min := targetSize * 2

	var dfs func(int, int, int)

	dfs = func(first, last, idx int) {
		if idx == originSize {
			temp := targetSize - match
			if min > temp {
				min = temp
			}
			return
		}

		for i := first; i <= last; i++ {
			if target[i] == origin[idx] {
				match++
			}

			dfs(i+1, last+1, idx+1)

			if target[i] == origin[idx] {
				match--
			}
		}
	}

	dfs(0, targetSize-originSize, 0)

	return min
}

func countDiff(bs1, bs2 []byte) int {
	res := 0
	for i := range bs1 {
		if bs1[i] != bs2[i] {
			res++
		}
	}

	return res
}
