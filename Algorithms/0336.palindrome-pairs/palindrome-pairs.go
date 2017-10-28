package Problem0336

func palindromePairs(words []string) [][]int {
	res := make([][]int, 0, len(words))
	size := len(words)
	if size < 2 {
		return res
	}

	hash := make(map[string]int, size)

	var i, j, jj int
	var ok bool

	for i = 0; i < size; i++ {
		hash[words[i]] = i
	}

	for i = 0; i < len(words); i++ {
		for j = 0; j <= len(words[i]); j++ {
			str1 := words[i][:j]
			str2 := words[i][j:]

			if isPalindrome(str2) {
				str1rvs := reverse(str1)
				if jj, ok = hash[str1rvs]; ok && jj != i {
					res = append(res, []int{i, jj})
				}
			}

			if isPalindrome(str1) {
				str2rvs := reverse(str2)
				if jj, ok = hash[str2rvs]; ok && jj != i && len(str1) != 0 {
					res = append(res, []int{jj, i})
				}
			}

		}
	}

	return res
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 反转字符串
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
