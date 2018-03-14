package problem0345

func reverseVowels(s string) string {
	bytes := []byte(s)
	i, j := 0, len(s)-1
	for {
		for i < len(s) && !isVowel(bytes[i]) {
			i++
		}
		for 0 <= j && !isVowel(bytes[j]) {
			j--
		}

		if i >= j {
			break
		}

		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
