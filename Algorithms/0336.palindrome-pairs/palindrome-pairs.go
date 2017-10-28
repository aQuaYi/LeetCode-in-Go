package Problem0336

func palindromePairs(words []string) [][]int {
	res := make([][]int, 0, len(words))

	var i, j int
	for i = 0; i < len(words); i++ {
		for j = i + 1; j < len(words); j++ {

			if isPalindrome(words[i] + words[j]) {
				res = append(res, []int{i, j})
			}

			if isPalindrome(words[j] + words[i]) {
				res = append(res, []int{j, i})
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
