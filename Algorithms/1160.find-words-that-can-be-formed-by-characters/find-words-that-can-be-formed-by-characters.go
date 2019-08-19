package problem1160

func countCharacters(words []string, chars string) int {
	char := parse(chars)
	res := 0
	for _, word := range words {
		res += count(char, word)
	}
	return res
}

func parse(s string) []int {
	res := make([]int, 26)
	for _, r := range s {
		res[r-'a']++
	}
	return res
}

func count(char []int, word string) int {
	res := 0
	w := make([]int, 26)
	for _, r := range word {
		b := r - 'a'
		w[b]++
		if w[b] > char[b] {
			return 0
		}
		res++
	}
	return res
}
