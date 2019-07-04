package problem1048

func longestStrChain(words []string) int {
	indexs := make([][]int, 17)
	count := make([]int, len(words))
	for i, word := range words {
		l := len(word)
		indexs[l] = append(indexs[l], i)
		count[i] = 1
	}

	res := 1
	for length := 1; length+1 <= 16; length++ {
		for _, i := range indexs[length] {
			for _, j := range indexs[length+1] {
				if count[j] > count[i] {
					// because of isPredecessor is expensive
					continue
				}
				if isPredecessor(words[i], words[j]) {
					count[j] = count[i] + 1
				}
			}
		}
	}

	for _, v := range count {
		res = max(res, v)
	}
	return res
}

func isPredecessor(w1, w2 string) bool {
	n := len(w1)
	diff := 0
	i, j := 0, 0
	for i < n && diff <= 1 {
		if w1[i] != w2[j] {
			diff++
		} else {
			i++
		}
		j++
	}
	return diff == 1 ||
		i == j // w1[:i]==w2[:i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
