package Problem0318

func maxProduct(words []string) int {
	max := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if !isShareLetters(words[i], words[j]) {
				if max < len(words[i])*len(words[j]) {
					max = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return max
}

func isShareLetters(a, b string) bool {
	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				return true
			}
		}
	}
	return false
}
