package Problem0438

func findAnagrams(s string, p string) []int {
	res := []int{}
	for i := 0; i+len(p) <= len(s); i++ {
		if isAnagram(p, s[i:i+len(p)]) {
			res = append(res, i)
		}
	}

	return res
}

func isAnagram(a, b string) bool {
	rec := [26]int{}
	for i := range a {
		rec[a[i]-'a']++
		rec[b[i]-'a']--
	}
	for _, r := range rec {
		if r != 0 {
			return false
		}
	}

	return true
}
