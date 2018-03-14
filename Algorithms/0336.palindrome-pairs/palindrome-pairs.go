package problem0336

func palindromePairs(words []string) [][]int {
	res := make([][]int, 0, len(words))
	size := len(words)
	if size < 2 {
		return res
	}

	hash := make(map[string]int, size)

	for i := 0; i < size; i++ {
		hash[words[i]] = i
	}

	for i := 0; i < len(words); i++ {
		for k := 0; k <= len(words[i]); k++ {
			right := words[i][k:]
			left := words[i][:k]

			if isPalindrome(right) {
				leftRev := reverse(left)
				if j, ok := hash[leftRev]; ok && i != j {
					res = append(res, []int{i, j})
				}
			}

			// 因为
			//   k == 0             的 right
			//   和
			//   k == len(words[i]) 的 left
			//   相等，会导致重复答案
			// 需要，len(left) != 0 来过滤掉这种情况
			if len(left) != 0 && isPalindrome(left) {
				rightRev := reverse(right)
				if j, ok := hash[rightRev]; ok && i != j {
					res = append(res, []int{j, i})
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
	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
