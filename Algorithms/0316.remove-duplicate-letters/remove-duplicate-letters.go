package Problem0316

func removeDuplicateLetters(s string) string {
	count := make([]int, 26)
	bytes := []byte(s)
	for _, b := range bytes {
		count[b-'a']++
	}

	var i, j, k int
	i, j = 0, 1

	for j < len(s) {
		if s[i] > s[j] {
			if count[s[i]-'a'] > 1 {
				bytes[i] = 0
				count[s[i]-'a']--
			}
			i++
			for i < len(s) && bytes[i] == 0 {
				i++
			}

			j++
			for j < len(s) && bytes[j] == 0 {
				j++
			}
		} else if s[i] < s[j] {
			if count[s[i]-'a'] > 1 {
				for k = j + 1; k < len(s); k++ {
					if bytes[k] == s[i] {
						bytes[k] = 0
						count[s[i]-'a']--
					}
				}
			}

			if count[s[j]-'a'] > 1 {
				for k = i - 1; 0 <= k; k-- {
					if bytes[k] == s[j] {
						bytes[k] = 0
						count[s[j]-'a']--
					}
				}
			}
			i++
			for i < len(s) && bytes[i] == 0 {
				i++
			}

			j++
			for j < len(s) && bytes[j] == 0 {
				j++
			}

		} else {
			bytes[j] = 0
			count[s[j]-'a']--
			j++
			for j < len(s) && bytes[j] == 0 {
				j++
			}
		}
	}

	res := ""
	for _, b := range bytes {
		if b != 0 {
			res += string(b)
		}
	}

	return res
}
