package Problem0383

func canConstruct(ransomNote string, magazine string) bool {
	rc := getCount(ransomNote)
	mc := getCount(magazine)

	for i, c := range rc {
		if mc[i] < c {
			return false
		}
	}
	return true
}

func getCount(s string) []int {
	res := make([]int, 26)
	for i := range s {
		res[s[i]-'a']++
	}
	return res
}
