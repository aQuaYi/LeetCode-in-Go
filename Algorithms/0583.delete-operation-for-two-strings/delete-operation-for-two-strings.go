package Problem0583

func minDistance(word1 string, word2 string) int {
	return len(word1) + len(word2) - comSub(word1, word2)*2
}

func comSub(word1, word2 string) int {
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}
	if len(word1) == 0 {
		return 0
	}

	res := 0

	for i := 0; i < len(word1) && len(word1)-i > res; i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				t, w1, w2 := getDiff(word1[i:], word2[j:])
				t += comSub(w1, w2)
				res = max(res, t)
			}
		}
	}

	return res
}

func getDiff(word1, word2 string) (int, string, string) {
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}

	var i int
	for i = 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			break
		}
	}

	return i, word1[i:], word2[i:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
