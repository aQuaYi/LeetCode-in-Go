package Problem0072

func minDistance(word1 string, word2 string) int {
	// 转换是可逆的
	// 互换可以只实现一套逻辑
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}

	size := len(word2)
	if size == 0 {
		return 0
	}

	origin := []byte(word1)
	target := []byte(word2)

	originSize := len(word1)
	if originSize == size {
		return countDiff(origin, target)
	}

	match := 0
	min := size * 2

	var dfs func(int, int, int)

	dfs = func(start, width, idx int) {
		if idx == originSize {
			temp := size - match
			if min > temp {
				min = temp
			}
			return
		}

		for i := start; i <= start+width; i++ {
			if origin[idx] == target[i] {
				match++
			}

			if width > 0 {
				dfs(i+1, width-1, idx+1)
			} else {
				dfs(i+1, 0, idx+1)
			}

			if origin[idx] == target[i] {
				match--
			}

		}
	}

	dfs(0, size-len(word1), 0)

	return min
}

func insert() {

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
