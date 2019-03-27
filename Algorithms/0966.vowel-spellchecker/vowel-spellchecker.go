package problem0966

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	isExactlyMatch := make(map[string]bool, len(wordlist))
	for _, w := range wordlist {
		isExactlyMatch[w] = true
	}

	cap := make(map[string]int, len(wordlist))
	vow := make(map[string]int, len(wordlist))
	for i := len(wordlist) - 1; i >= 0; i-- {
		w := wordlist[i]
		cap[strings.ToLower(w)] = i
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

func check(q string, wordlist []string) string {
	for _, w := range wordlist {
		if strings.ToLower(q) == strings.ToLower(w) {
			return w
		}
	}
	for _, w := range wordlist {
		if replacingVowel(q) == replacingVowel(w) {
			return w
		}
	}
	return ""
}

func replacingVowel(s string) string {
	vowels := []string{"a", "e", "i", "o", "u"}
	s = strings.ToLower(s)
	for _, v := range vowels {
		s = strings.Replace(s, v, "*", -1)
	}
	return s
}
