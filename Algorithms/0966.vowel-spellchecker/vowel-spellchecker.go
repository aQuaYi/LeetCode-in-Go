package problem0966

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	isExactlyMatch := make(map[string]bool, len(wordlist))
	cap := make(map[string]int, len(wordlist))
	vow := make(map[string]int, len(wordlist))
	for i := len(wordlist) - 1; i >= 0; i-- {
		w := wordlist[i]
		isExactlyMatch[w] = true // case sensitive
		w = strings.ToLower(w)
		cap[w] = i // case insensitive
		vow[replacingVowel(w)] = i
	}

	corrects := make([]string, len(queries))
	for i, q := range queries {
		if isExactlyMatch[q] {
			corrects[i] = q
		} else if j, ok := cap[strings.ToLower(q)]; ok {
			corrects[i] = wordlist[j]
		} else if k, ok := vow[replacingVowel(q)]; ok {
			corrects[i] = wordlist[k]
		}
	}

	return corrects
}

func replacingVowel(s string) string {
	s = strings.ToLower(s) // case insensitive
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, v := range vowels {
		s = strings.Replace(s, v, "*", -1)
	}
	return s
}
